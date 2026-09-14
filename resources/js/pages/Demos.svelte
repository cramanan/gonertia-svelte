<script module lang="ts">
  export { default as layout } from "@/layouts/AppLayout.svelte";
</script>

<script lang="ts">
  import { Link } from "@inertiajs/svelte";
  import AppHead from "@/components/AppHead.svelte";

  const demos = [
    {
      title: "Error Pages",
      description: "Errors renders in a special page with the right HTTP status",
      href: "/non-existent",
    },
    {
      title: "Static Asset Handling",
      description: "...",
      href: "/demos/assets",
    },
    {
      title: "API Fetching",
      description: "Fetch JSON from a Go endpoint and render the response.",
      href: "/demos/api-fetching",
    },
  ];
</script>

<AppHead title="Demos" />

<div class="mx-auto max-w-2xl px-6 py-12">
  <h1 class="mb-1 font-medium text-lg">Demos</h1>
  <p class="mb-6 text-[#706f6c] dark:text-[#A1A09A]">
    Small examples showing off features of this stack.
  </p>

  <ul class="flex flex-col divide-y divide-[#e3e3e0] dark:divide-[#3E3E3A]">
    {#each demos as demo (demo.href)}
      <li>
        <Link
          href={demo.href}
          class="block py-4 hover:text-[#00ADD8] dark:hover:text-[#5DC9E2]"
        >
          <span class="font-medium">{demo.title}</span>
          <p class="text-sm text-[#706f6c] dark:text-[#A1A09A]">
            {demo.description}
          </p>
        </Link>
      </li>
    {/each}
  </ul>
</div>
