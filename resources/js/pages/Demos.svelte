<script module lang="ts">
  export { default as layout } from "@/layouts/AppLayout.svelte";
</script>

<script lang="ts">
  import { Link } from "@inertiajs/svelte";
  import AppHead from "@/components/AppHead.svelte";
  import * as Card from "@/components/ui/card";

  const demos = [
    {
      title: "Error Pages",
      description: "Errors renders in a special page with the right HTTP status",
      href: "/non-existent",
    },
    {
      title: "Static Asset Handling",
      description: "Vite-fingerprinted imports next to files served from public/.",
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
  <p class="mb-6 text-muted-foreground">
    Small examples showing off features of this stack.
  </p>

  <ul class="flex flex-col gap-3">
    {#each demos as demo (demo.href)}
      <li>
        <Link href={demo.href} class="block">
          <Card.Root class="transition-colors hover:bg-accent">
            <Card.Header>
              <Card.Title>{demo.title}</Card.Title>
              <Card.Description>{demo.description}</Card.Description>
            </Card.Header>
          </Card.Root>
        </Link>
      </li>
    {/each}
  </ul>
</div>
