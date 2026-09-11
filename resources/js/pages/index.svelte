<script module lang="ts">
  // import Layout from "@/layouts/layout.svelte";
  // export const layout = Layout;
  export { default as layout } from "@/layouts/layout.svelte";
</script>

<script lang="ts">
  let count = $state(0);
  let status = $state();

  $effect(
    () =>
      void fetch("/api/status")
        .then((response) => response.json())
        .then((value) => (status = value)),
  );
</script>

<h1>Hello World</h1>

<div>
  <button onclick={() => count++}>{count}</button>
  <code style="white-space: pre-wrap;">
    {JSON.stringify(status, null, "\t")}
  </code>
</div>
