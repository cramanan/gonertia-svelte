<script module lang="ts">
  export { default as layout } from "@/layouts/AppLayout.svelte";
</script>

<script lang="ts">
  import AppHead from "@/components/AppHead.svelte";

  let status = $state<Record<string, string>>();

  $effect(
    () =>
      void fetch("/api/status")
        .then((response) => response.json())
        .then((value) => (status = value)),
  );
</script>

<AppHead title="API Fetching" />

<div class="mx-auto max-w-2xl px-6 py-12">
  <h1 class="mb-1 font-medium text-lg">API Fetching</h1>
  <p class="mb-6 text-[#706f6c] dark:text-[#A1A09A]">
    Fetches <code>/api/status</code> from the Go backend on mount.
  </p>

  <pre
    class="rounded-lg bg-white p-4 text-[13px] shadow-[inset_0px_0px_0px_1px_rgba(26,26,0,0.16)] dark:bg-[#161615] dark:shadow-[inset_0px_0px_0px_1px_#fffaed2d]"><code>{JSON.stringify(status, null, 2)}</code></pre>
</div>
