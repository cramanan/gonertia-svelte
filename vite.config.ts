import fs from "node:fs";
import { defineConfig, lazyPlugins, loadEnv, type Plugin } from "vite-plus";
import laravel from "laravel-vite-plugin";
import inertia from "@inertiajs/vite";
import tailwindcss from "@tailwindcss/vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// Serves our own page when the Vite dev server is opened directly. Registered
// as a pre-middleware, it answers before laravel-vite-plugin's post-middleware
// would serve its Laravel-branded one.
function devServerPage(file: string): Plugin {
  return {
    name: "gravel:dev-server-page",
    apply: "serve",
    configureServer(server) {
      const envDir = server.config.envDir || process.cwd();
      const appUrl = loadEnv(server.config.mode, envDir, "APP_URL").APP_URL ?? "undefined";
      server.middlewares.use((req, res, next) => {
        if (req.url !== "/" && req.url !== "/index.html") return next();
        res.statusCode = 404;
        res.setHeader("Content-Type", "text/html");
        res.end(fs.readFileSync(file, "utf8").replace(/{{ APP_URL }}/g, appUrl));
      });
    },
  };
}

export default defineConfig({
  staged: {
    "*": "vp check --fix",
  },
  fmt: {
    // Go html/template: the formatter splits {{ ... }} actions.
    ignorePatterns: ["resources/views/**"],
  },
  lint: {
    jsPlugins: [{ name: "vite-plus", specifier: "vite-plus/oxlint-plugin" }],
    rules: {
      "vite-plus/prefer-vite-plus-imports": "error",
    },
    // typeCheck can't see `<script module>` exports from .svelte; svelte-check (`check` script) covers types.
    options: { typeAware: true },
  },
  plugins: lazyPlugins(() => [
    laravel({
      input: ["resources/css/app.css", "resources/js/app.ts"],
      refresh: true,
    }),
    devServerPage("resources/views/dev-server.html"),
    inertia(),
    tailwindcss(),
    svelte(),
  ]),
});
