<script lang="ts">
  import NavBar from "@/components/NavBar.svelte";
  import type { Snippet } from "svelte";

  let { children } : { children : Snippet } = $props(); 
</script>

<div class="flex min-h-screen flex-col bg-[#FDFDFC] text-[#1b1b18] dark:bg-[#0a0a0a] dark:text-[#EDEDEC]">
  <header class="border-b border-[#e3e3e0] dark:border-[#3E3E3A]">
    <NavBar />
  </header>
  <main class="flex flex-1 flex-col">
    {@render children()}
  </main>
  <footer
    class="border-t border-[#e3e3e0] px-6 py-6 text-center text-sm text-[#706f6c] dark:border-[#3E3E3A] dark:text-[#A1A09A]"
  >
    <p>
      &copy; {new Date().getFullYear()} Gravel. Built with
      <a
        href="https://go.dev"
        target="_blank"
        class="font-medium text-[#00ADD8] underline underline-offset-4 dark:text-[#5DC9E2]"
      >
        Go
      </a>
      &
      <a
        href="https://inertiajs.com"
        target="_blank"
        class="font-medium text-[#00ADD8] underline underline-offset-4 dark:text-[#5DC9E2]"
      >
        Inertia
      </a>.
    </p>
  </footer>
</div>