<script module lang="ts">
  export { default as layout } from "@/layouts/AppLayout.svelte";
</script>

<script lang="ts">
  import AppHead from "@/components/AppHead.svelte";
  import inertiaLogo from "@/assets/inertia-logo.svg";

  const assets = [
    {
      title: "Imported by Vite",
      description:
        "Imported from resources/js/assets. Vite fingerprints it into public/build/assets; the URL changes when the file does.",
      src: inertiaLogo,
    },
    {
      title: "Served from public/",
      description:
        "Copied as-is from public/ and embedded in the Go binary, served by http.FileServerFS at a fixed URL.",
      src: "/go-logo-blue.svg",
    },
  ];
</script>

<AppHead title="Static Asset Handling" />

<div class="mx-auto max-w-2xl px-6 py-12">
  <h1 class="mb-1 font-medium text-lg">Static Asset Handling</h1>
  <p class="mb-6 text-[#706f6c] dark:text-[#A1A09A]">
    Two ways to ship an asset with this stack.
  </p>

  <ul class="flex flex-col gap-4">
    {#each assets as asset (asset.src)}
      <li
        class="flex items-center gap-6 rounded-lg bg-white p-4 shadow-[inset_0px_0px_0px_1px_rgba(26,26,0,0.16)] dark:bg-[#161615] dark:shadow-[inset_0px_0px_0px_1px_#fffaed2d]"
      >
        <img src={asset.src} alt="" class="h-12 w-24 shrink-0 object-contain" />
        <div class="min-w-0">
          <span class="font-medium">{asset.title}</span>
          <p class="text-sm text-[#706f6c] dark:text-[#A1A09A]">{asset.description}</p>
          <code class="block truncate text-[13px]">{asset.src}</code>
        </div>
      </li>
    {/each}
  </ul>
</div>
