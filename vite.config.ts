import { defineConfig } from "vite";
import gravel from "./gravel-vite-plugin/index.ts";
import inertia from "@inertiajs/vite";
import tailwindcss from "@tailwindcss/vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

export default defineConfig({
  plugins: [
    gravel({
      input: ["resources/css/app.css", "resources/js/app.ts"],
      refresh: true,
    }),
    inertia(),
    tailwindcss(),
    svelte(),
  ],
});
