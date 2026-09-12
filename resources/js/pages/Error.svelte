<script module lang="ts">
  export { default as layout } from "@/layouts/AppLayout.svelte";
</script>

<script lang="ts">
  import { Link } from "@inertiajs/svelte";
  import AppHead from "@/components/AppHead.svelte";

  let { status }: { status: number } = $props();

  const messages: Record<number, { title: string; description: string }> = {
    403: {
      title: "Forbidden",
      description: "You don't have permission to access this page.",
    },
    404: {
      title: "Not Found",
      description: "The page you're looking for doesn't exist.",
    },
    500: {
      title: "Server Error",
      description: "Something went wrong on our end.",
    },
    503: {
      title: "Unavailable",
      description: "The service is temporarily down for maintenance.",
    },
  };

  let message = $derived(
    messages[status] ?? {
      title: "Error",
      description: "An unexpected error occurred.",
    },
  );
</script>

<AppHead title={message.title} />

<div class="m-auto flex max-w-2xl flex-col items-center px-6 py-24 text-center">
  <p class="mb-2 font-mono text-5xl font-medium text-[#00ADD8] dark:text-[#5DC9E2]">
    {status}
  </p>
  <h1 class="mb-2 font-medium text-lg">{message.title}</h1>
  <p class="mb-6 text-[#706f6c] dark:text-[#A1A09A]">
    {message.description}
  </p>
  <Link
    href="/"
    class="text-sm text-[#00ADD8] hover:underline dark:text-[#5DC9E2]"
  >
    Back home
  </Link>
</div>
