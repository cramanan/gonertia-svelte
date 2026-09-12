<script>
  import { Link } from "@inertiajs/svelte";
  import inertiaLogo from "@/assets/inertia-logo.svg";
</script>

<nav class="flex items-center gap-6 px-6 py-4">
  <Link href="/" class="flex items-center gap-2" aria-label="Home">
    <img src="/go-logo-blue.svg" alt="Go" class="h-5 w-auto" />
    <img src={inertiaLogo} alt="Inertia" height="16" class="h-4 w-auto" />
  </Link>
  <Link
    href="/demos"
    class="text-sm font-medium text-[#1b1b18] hover:text-[#00ADD8] dark:text-[#EDEDEC] dark:hover:text-[#5DC9E2]"
  >
    Demos
  </Link>
</nav>