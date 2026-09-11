import inertia from "@inertiajs/vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import { defineConfig, type Plugin, type Rolldown } from "vite";
import fs from "node:fs";
import path from "node:path";

interface PluginConfig {
  /**
   * The path or paths of the entry points to compile.
   */
  input: Rolldown.InputOption;
}

interface GravelPlugin extends Plugin {}

function gravel(pluginConfig: PluginConfig): GravelPlugin {
  const hotFile = "public/hot";
  const cleanup = () => {
    try {
      fs.unlinkSync(hotFile);
    } catch {
      // already gone
    }
  };

  return {
    name: "gravel",
    config(config, { command }) {
      return {
        base: command === "build" ? "/build/" : undefined,
        build: {
          outDir: config.build?.outDir,
          manifest: config.build?.manifest ?? "manifest.json",
          rolldownOptions: {
            input:
              config.build?.rolldownOptions?.input ??
              config.build?.rollupOptions?.input ??
              pluginConfig.input,
          },
        },
      };
    },
    configureServer(server) {
      server.httpServer?.once("listening", () => {
        const address = server.httpServer?.address();
        if (address && typeof address !== "string") {
          fs.writeFileSync(hotFile, `http://localhost:${address.port}`);
        }
      });
      process.once("exit", cleanup);
      process.once("SIGINT", () => {
        cleanup();
        process.exit();
      });
      process.once("SIGTERM", () => {
        cleanup();
        process.exit();
      });
    },
    buildEnd: cleanup,
  };
}
// https://vite.dev/config/
export default defineConfig({
  plugins: [
    gravel({ input: ["resources/css/app.css", "resources/js/app.ts"] }),
    inertia(),
    svelte(),
  ],
  resolve: {
    alias: { "@": path.resolve(import.meta.dirname, "resources/js") },
  },
  build: { outDir: "vite/dist" },
});
