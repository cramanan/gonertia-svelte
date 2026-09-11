<script module lang="ts">
  // import Layout from "@/layouts/layout.svelte";
  // export const layout = Layout;
  export { default as layout } from "@/layouts/layout.svelte";

  let count = $state(0);
</script>

Hello World

<button onclick={() => count++}>{count}</button>
