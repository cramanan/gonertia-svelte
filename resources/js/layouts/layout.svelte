<script lang="ts">
  import type { Snippet } from "svelte";

  const { children }: { children: Snippet } = $props();
</script>

<header>HEADER</header>
<main>{@render children()}</main>
<footer>FOOTER</footer>
