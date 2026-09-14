import fs, { globSync } from "node:fs";
import path from "path";
import {
  createLogger,
  defaultAllowedOrigins,
  loadEnv,
  type ConfigEnv,
  type Plugin,
  type PluginOption,
  type ResolvedConfig,
  type Rolldown,
  type UserConfig,
} from "vite";
import type { AddressInfo } from "node:net";
import colors from "picocolors";
import { fileURLToPath } from "node:url";
import fullReload, {
  type Config as FullReloadConfig,
} from "vite-plugin-full-reload";

interface PluginConfig {
  /**
   * The path or paths of the entry points to compile.
   */
  input: Rolldown.InputOption;

  /**
   * Laravel's public directory.
   *
   * @default 'public'
   */
  publicDirectory?: string;

  /**
   * The public subdirectory where compiled assets should be written.
   *
   * @default 'build'
   */
  buildDirectory?: string;

  /**
   * The path to the "hot" file.
   *
   * @default `${publicDirectory}/hot`
   */
  hotFile?: string;

  /**
   * Configuration for performing full page refresh on blade (or other) file changes.
   *
   * {@link https://github.com/ElMassimo/vite-plugin-full-reload}
   * @default false
   */
  refresh?: boolean | string | string[] | RefreshConfig | RefreshConfig[];

  /**
   * Transform the code while serving.
   */
  transformOnServe?: (code: string, url: DevServerUrl) => string;

  /**
   * Asset file glob patterns to include in the build.
   *
   * Files matching these patterns will be processed and versioned by Vite,
   * even if they are not imported in your JavaScript. This is useful for
   * assets referenced in Blade templates via `Vite::asset()`.
   *
   * @default []
   */
  assets?: string | string[];
}

interface RefreshConfig {
  paths: string[];
  config?: FullReloadConfig;
}

interface GravelPlugin extends Plugin {
  config: (config: UserConfig, env: ConfigEnv) => UserConfig;
}

type DevServerUrl = `${"http" | "https"}://${string}:${number}`;

const DEV_SERVER_ORIGIN_PLACEHOLDER = "http://__gravel_vite_placeholder__.test";

let exitHandlersBound = false;

export const refreshPaths = ["resources/views/**"].filter((path) =>
  fs.existsSync(path.replace(/\*\*$/, "")),
);

const logger = createLogger("info", { prefix: "[gravel-vite-plugin]" });

/**
 * The directory of the current file.
 */
function dirname(): string {
  return fileURLToPath(new URL(".", import.meta.url));
}

/**
 * Convert the users configuration into a standard structure with defaults.
 */
function resolvePluginConfig(
  config: string | string[] | PluginConfig,
): Required<PluginConfig> {
  if (typeof config === "undefined") {
    throw new Error("gravel-vite-plugin: missing configuration.");
  }

  if (typeof config === "string" || Array.isArray(config)) {
    config = { input: config };
  }

  if (typeof config.input === "undefined") {
    throw new Error('gravel-vite-plugin: missing configuration for "input".');
  }

  if (typeof config.publicDirectory === "string") {
    config.publicDirectory = config.publicDirectory.trim().replace(/^\/+/, "");

    if (config.publicDirectory === "") {
      throw new Error(
        "laravel-vite-plugin: publicDirectory must be a subdirectory. E.g. 'public'.",
      );
    }
  }

  if (typeof config.buildDirectory === "string") {
    config.buildDirectory = config.buildDirectory
      .trim()
      .replace(/^\/+/, "")
      .replace(/\/+$/, "");

    if (config.buildDirectory === "") {
      throw new Error(
        "laravel-vite-plugin: buildDirectory must be a subdirectory. E.g. 'build'.",
      );
    }
  }

  if (config.refresh === true) {
    config.refresh = [{ paths: refreshPaths }];
  }

  return {
    input: config.input,
    publicDirectory: config.publicDirectory ?? "public",
    buildDirectory: config.buildDirectory ?? "build",
    refresh: config.refresh ?? false,
    hotFile:
      config.hotFile ?? path.join(config.publicDirectory ?? "public", "hot"),
    transformOnServe: config.transformOnServe ?? ((code) => code),
    assets:
      typeof config.assets === "string"
        ? [config.assets]
        : (config.assets ?? []),
  };
}

/**
 * Laravel plugin for Vite.
 *
 * @param config - A config object or relative path(s) of the scripts to be compiled.
 */
export default function gravel(
  config: string | string[] | PluginConfig,
): [GravelPlugin, ...Plugin[]] {
  const pluginConfig = resolvePluginConfig(config);

  return [
    resolveGravelPlugin(pluginConfig),
    ...resolveAssetPlugin(pluginConfig.assets),
    ...(resolveFullReloadConfig(pluginConfig) as Plugin[]),
  ];
}

function resolveGravelPlugin(
  pluginConfig: Required<PluginConfig>,
): GravelPlugin {
  let viteDevServerUrl: DevServerUrl;
  let resolvedConfig: ResolvedConfig;
  let userConfig: UserConfig;

  const defaultAliases: Record<string, string> = {
    "@": "/resources/js",
  };

  const removeHotFile = () => {
    try {
      fs.rmSync(pluginConfig.hotFile);
    } catch {}
  };

  return {
    name: "gravel",
    enforce: "post",
    config(config, { command, mode }) {
      userConfig = config;
      const env = loadEnv(mode, userConfig.envDir || process.cwd(), "");
      const assetUrl = env.ASSET_URL ?? "";
      return {
        base:
          userConfig.base ??
          (command === "build" ? resolveBase(pluginConfig, assetUrl) : ""),
        publicDir: userConfig.publicDir ?? false,
        build: {
          manifest: userConfig.build?.manifest ?? "manifest.json",
          outDir: config.build?.outDir ?? resolveOutDir(pluginConfig),
          rolldownOptions: {
            input:
              userConfig.build?.rolldownOptions?.input ??
              userConfig.build?.rollupOptions?.input ??
              resolveInput(pluginConfig),
          },
          assetsInlineLimit: userConfig.build?.assetsInlineLimit ?? 0,
        },
        server: {
          origin:
            userConfig.server?.origin ??
            (command === "serve" ? DEV_SERVER_ORIGIN_PLACEHOLDER : undefined),
          cors: userConfig.server?.cors ?? {
            origin: userConfig.server?.origin ?? [
              defaultAllowedOrigins,
              ...(env.APP_URL ? [env.APP_URL] : []), // *               (APP_URL="http://my-app.tld")
              // /^https?:\/\/.*\.test(:\d+)?$/, // Valet / Herd    (SCHEME://*.test:PORT)
            ],
          },
        },
        resolve: {
          alias: Array.isArray(userConfig.resolve?.alias)
            ? [
                ...(userConfig.resolve?.alias ?? []),
                ...Object.keys(defaultAliases).map((alias) => ({
                  find: alias,
                  replacement: defaultAliases[alias],
                })),
              ]
            : {
                ...defaultAliases,
                ...userConfig.resolve?.alias,
              },
        },
      };
    },

    configResolved(config) {
      resolvedConfig = config;
    },

    buildStart: removeHotFile,

    buildEnd: removeHotFile,

    transform(code) {
      if (resolvedConfig.command === "serve" && viteDevServerUrl) {
        return code.replaceAll(DEV_SERVER_ORIGIN_PLACEHOLDER, viteDevServerUrl);
      }
    },

    configureServer(server) {
      if (process.env.VITEST !== undefined) {
        return;
      }

      const envDir = resolvedConfig.envDir || process.cwd();
      const appUrl =
        loadEnv(resolvedConfig.mode, envDir, "APP_URL").APP_URL ?? "undefined";

      server.httpServer?.once("listening", () => {
        const address = server.httpServer?.address();
        const isAddressInfo = (
          x: string | AddressInfo | null | undefined,
        ): x is AddressInfo => typeof x === "object" && x !== null;
        if (isAddressInfo(address)) {
          viteDevServerUrl =
            userConfig.server?.origin &&
            userConfig.server.origin !== DEV_SERVER_ORIGIN_PLACEHOLDER
              ? (userConfig.server.origin as DevServerUrl)
              : resolveDevServerUrl(address, server.config);

          const hotFileParentDirectory = path.dirname(pluginConfig.hotFile);

          if (!fs.existsSync(hotFileParentDirectory)) {
            fs.mkdirSync(hotFileParentDirectory, { recursive: true });
            setTimeout(() => {
              logger.info(
                `Hot file directory created ${colors.dim(fs.realpathSync(hotFileParentDirectory))}`,
                { clear: true, timestamp: true },
              );
            }, 200);
          }

          fs.writeFileSync(
            pluginConfig.hotFile,
            `${viteDevServerUrl}${server.config.base.replace(/\/$/, "")}`,
          );
        }
      });

      if (!exitHandlersBound) {
        const clean = () => {
          if (fs.existsSync(pluginConfig.hotFile)) {
            fs.rmSync(pluginConfig.hotFile);
          }
        };

        process.on("exit", clean);
        process.on("SIGINT", () => process.exit());
        process.on("SIGTERM", () => process.exit());
        process.on("SIGHUP", () => process.exit());

        exitHandlersBound = true;
      }

      return () =>
        server.middlewares.use((req, res, next) => {
          if (req.url === "/index.html") {
            res.statusCode = 404;

            res.end(
              fs
                .readFileSync(path.join(dirname(), "dev-server-index.html"))
                .toString()
                .replace(/{{ APP_URL }}/g, appUrl),
            );
          }

          next();
        });
    },
  };
}

function resolveFullReloadConfig({
  refresh: config,
}: Required<PluginConfig>): PluginOption[] {
  if (typeof config === "boolean") {
    return [];
  }

  if (typeof config === "string") {
    config = [{ paths: [config] }];
  }

  if (!Array.isArray(config)) {
    config = [config];
  }

  if (config.some((c) => typeof c === "string")) {
    config = [{ paths: config }] as RefreshConfig[];
  }

  return (config as RefreshConfig[]).flatMap((c) =>
    fullReload(c.paths, c.config),
  );
}

/**
 * Resolve the asset-emitting plugin from the configuration.
 */
function resolveAssetPlugin(assets: string | string[]): Plugin[] {
  if (assets.length === 0) {
    return [];
  }

  return [
    {
      name: "laravel:assets",
      apply: "build",
      buildStart() {
        for (const file of globSync(assets)) {
          if (fs.statSync(file).isFile()) {
            this.emitFile({
              type: "asset",
              name: path.basename(file),
              originalFileName: file,
              source: fs.readFileSync(file),
            });
          }
        }
      },
    },
  ];
}

/**
 * Resolve the Vite input path from the configuration.
 */
function resolveInput(
  config: Required<PluginConfig>,
): Rolldown.InputOption | undefined {
  return config.input;
}

/**
 * Resolve the Vite base option from the configuration.
 */
function resolveBase(config: Required<PluginConfig>, assetUrl: string): string {
  return (
    assetUrl +
    (!assetUrl.endsWith("/") ? "/" : "") +
    config.buildDirectory +
    "/"
  );
}

/**
 * Resolve the Vite outDir path from the configuration.
 */
function resolveOutDir(config: Required<PluginConfig>) {
  return path.join(config.publicDirectory, config.buildDirectory);
}

function resolveDevServerUrl(
  address: AddressInfo,
  config: ResolvedConfig,
): DevServerUrl {
  const configHmrProtocol =
    typeof config.server.hmr === "object" ? config.server.hmr.protocol : null;
  const clientProtocol = configHmrProtocol
    ? configHmrProtocol === "wss"
      ? "https"
      : "http"
    : null;
  const serverProtocol = config.server.https ? "https" : "http";
  const protocol = clientProtocol ?? serverProtocol;

  const configHmrHost =
    typeof config.server.hmr === "object" ? config.server.hmr.host : null;
  const configHost =
    typeof config.server.host === "string" ? config.server.host : null;
  const serverAddress = isIpv6(address)
    ? `[${address.address}]`
    : address.address;
  const host = configHmrHost ?? configHost ?? serverAddress;

  const configHmrClientPort =
    typeof config.server.hmr === "object" ? config.server.hmr.clientPort : null;
  const port = configHmrClientPort ?? address.port;

  return `${protocol}://${host}:${port}`;
}

function isIpv6(address: AddressInfo): boolean {
  return (
    address.family === "IPv6" ||
    // In node >=18.0 <18.4 this was an integer value. This was changed in a minor version.
    // See: https://github.com/laravel/vite-plugin/issues/103
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore-next-line
    address.family === 6
  );
}
