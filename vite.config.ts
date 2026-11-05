import { defineConfig, lazyPlugins } from "vite-plus";
import gravel from "./gravel-vite-plugin/index.ts";
import inertia from "@inertiajs/vite";
import tailwindcss from "@tailwindcss/vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

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
    gravel({
      input: ["resources/css/app.css", "resources/js/app.ts"],
      refresh: true,
    }),
    inertia(),
    tailwindcss(),
    svelte(),
  ]),
});
