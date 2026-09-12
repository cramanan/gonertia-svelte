<script lang="ts">
  import { Link } from "@inertiajs/svelte";
  import AppHead from "@/components/AppHead.svelte";
</script>

<AppHead title="Welcome">
  <link rel="preconnect" href="https://rsms.me/" />
  <link rel="stylesheet" href="https://rsms.me/inter/inter.css" />
</AppHead>

<div
  class="flex min-h-screen flex-col items-center bg-[#FDFDFC] p-6 text-[#1b1b18] lg:justify-center lg:p-8 dark:bg-[#0a0a0a]"
>
  <div
    class="flex w-full items-center justify-center opacity-100 transition-opacity duration-750 lg:grow starting:opacity-0"
  >
    <main
      class="flex w-full max-w-83.75 flex-col-reverse overflow-hidden rounded-lg lg:max-w-4xl lg:flex-row"
    >
      <div
        class="flex-1 rounded-br-lg rounded-bl-lg bg-white p-6 pb-12 text-[13px] leading-5 shadow-[inset_0px_0px_0px_1px_rgba(26,26,0,0.16)] lg:rounded-tl-lg lg:rounded-br-none lg:p-20 dark:bg-[#161615] dark:text-[#EDEDEC] dark:shadow-[inset_0px_0px_0px_1px_#fffaed2d]"
      >
        <h1 class="mb-1 font-medium">Let's get started</h1>
        <p class="mb-2 text-[#706f6c] dark:text-[#A1A09A]">
          Go has a simple, fast, and reliable ecosystem. <br />We suggest starting
          with the following.
        </p>
        <ul class="mb-4 flex flex-col lg:mb-6">
          <li
            class="relative flex items-center gap-4 py-2 before:absolute before:top-1/2 before:bottom-0 before:left-[0.4rem] before:border-l before:border-[#e3e3e0] dark:before:border-[#3E3E3A]"
          >
            <span class="relative bg-white py-1 dark:bg-[#161615]">
              <span
                class="flex h-3.5 w-3.5 items-center justify-center rounded-full border border-[#e3e3e0] bg-[#FDFDFC] shadow-[0px_0px_1px_0px_rgba(0,0,0,0.03),0px_1px_2px_0px_rgba(0,0,0,0.06)] dark:border-[#3E3E3A] dark:bg-[#161615]"
              >
                <span
                  class="h-1.5 w-1.5 rounded-full bg-[#dbdbd7] dark:bg-[#3E3E3A]"
                ></span>
              </span>
            </span>
            <span>
              Read the
              <a
                href="https://go.dev/doc/"
                target="_blank"
                class="ml-1 inline-flex items-center space-x-1 font-medium text-[#00ADD8] underline underline-offset-4 dark:text-[#5DC9E2]"
              >
                <span>Documentation</span>
                <svg
                  width="10"
                  height="11"
                  viewBox="0 0 10 11"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-2.5 w-2.5"
                >
                  <path
                    d="M7.70833 6.95834V2.79167H3.54167M2.5 8L7.5 3.00001"
                    stroke="currentColor"
                    stroke-linecap="square"
                  />
                </svg>
              </a>
            </span>
          </li>
          <li
            class="relative flex items-center gap-4 py-2 before:absolute before:top-0 before:bottom-1/2 before:left-[0.4rem] before:border-l before:border-[#e3e3e0] dark:before:border-[#3E3E3A]"
          >
            <span class="relative bg-white py-1 dark:bg-[#161615]">
              <span
                class="flex h-3.5 w-3.5 items-center justify-center rounded-full border border-[#e3e3e0] bg-[#FDFDFC] shadow-[0px_0px_1px_0px_rgba(0,0,0,0.03),0px_1px_2px_0px_rgba(0,0,0,0.06)] dark:border-[#3E3E3A] dark:bg-[#161615]"
              >
                <span
                  class="h-1.5 w-1.5 rounded-full bg-[#dbdbd7] dark:bg-[#3E3E3A]"
                ></span>
              </span>
            </span>
            <span>
              Learn by example at
              <a
                href="https://gobyexample.com"
                target="_blank"
                class="ml-1 inline-flex items-center space-x-1 font-medium text-[#00ADD8] underline underline-offset-4 dark:text-[#5DC9E2]"
              >
                <span>Go by Example</span>
                <svg
                  width="10"
                  height="11"
                  viewBox="0 0 10 11"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-2.5 w-2.5"
                >
                  <path
                    d="M7.70833 6.95834V2.79167H3.54167M2.5 8L7.5 3.00001"
                    stroke="currentColor"
                    stroke-linecap="square"
                  />
                </svg>
              </a>
            </span>
          </li>
        </ul>
        <ul class="flex gap-3 text-sm leading-normal">
          <li>
            <a
              href="https://go.dev/solutions/cloud"
              target="_blank"
              class="inline-block rounded-sm border border-black bg-[#1b1b18] px-5 py-1.5 text-sm leading-normal text-white hover:border-black hover:bg-black dark:border-[#eeeeec] dark:bg-[#eeeeec] dark:text-[#1C1C1A] dark:hover:border-white dark:hover:bg-white"
            >
              Deploy now
            </a>
          </li>
        </ul>
      </div>
      <div
        class="bg-[#E6FBFF] dark:bg-[#001B22] relative lg:-ml-px -mb-px lg:mb-0 rounded-t-lg lg:rounded-t-none lg:rounded-r-lg aspect-335/364 lg:aspect-auto w-full lg:w-109.5 shrink-0 overflow-hidden flex items-center justify-center"
      >
        <!-- Go Logo -->
        <svg
          height="140"
          viewBox="0 0 207 78"
          xmlns="http://www.w3.org/2000/svg"
          class="w-auto max-w-[80%] text-[#00ADD8] dark:text-[#5DC9E2] transition-all translate-y-0 opacity-100 duration-750 starting:opacity-0 motion-safe:starting:translate-y-6"
        >
          <g fill="currentColor" fill-rule="evenodd">
            <path
              d="m16.2 24.1c-.4 0-.5-.2-.3-.5l2.1-2.7c.2-.3.7-.5 1.1-.5h35.7c.4 0 .5.3.3.6l-1.7 2.6c-.2.3-.7.6-1 .6z"
            />
            <path
              d="m1.1 33.3c-.4 0-.5-.2-.3-.5l2.1-2.7c.2-.3.7-.5 1.1-.5h45.6c.4 0 .6.3.5.6l-.8 2.4c-.1.4-.5.6-.9.6z"
            />
            <path
              d="m25.3 42.5c-.4 0-.5-.3-.3-.6l1.4-2.5c.2-.3.6-.6 1-.6h20c.4 0 .6.3.6.7l-.2 2.4c0 .4-.4.7-.7.7z"
            />
            <g transform="translate(55)">
              <path
                d="m74.1 22.3c-6.3 1.6-10.6 2.8-16.8 4.4-1.5.4-1.6.5-2.9-1-1.5-1.7-2.6-2.8-4.7-3.8-6.3-3.1-12.4-2.2-18.1 1.5-6.8 4.4-10.3 10.9-10.2 19 .1 8 5.6 14.6 13.5 15.7 6.8.9 12.5-1.5 17-6.6.9-1.1 1.7-2.3 2.7-3.7-3.6 0-8.1 0-19.3 0-2.1 0-2.6-1.3-1.9-3 1.3-3.1 3.7-8.3 5.1-10.9.3-.6 1-1.6 2.5-1.6h36.4c-.2 2.7-.2 5.4-.6 8.1-1.1 7.2-3.8 13.8-8.2 19.6-7.2 9.5-16.6 15.4-28.5 17-9.8 1.3-18.9-.6-26.9-6.6-7.4-5.6-11.6-13-12.7-22.2-1.3-10.9 1.9-20.7 8.5-29.3 7.1-9.3 16.5-15.2 28-17.3 9.4-1.7 18.4-.6 26.5 4.9 5.3 3.5 9.1 8.3 11.6 14.1.6.9.2 1.4-1 1.7z"
              />
              <path
                d="m107.2 77.6c-9.1-.2-17.4-2.8-24.4-8.8-5.9-5.1-9.6-11.6-10.8-19.3-1.8-11.3 1.3-21.3 8.1-30.2 7.3-9.6 16.1-14.6 28-16.7 10.2-1.8 19.8-.8 28.5 5.1 7.9 5.4 12.8 12.7 14.1 22.3 1.7 13.5-2.2 24.5-11.5 33.9-6.6 6.7-14.7 10.9-24 12.8-2.7.5-5.4.6-8 .9zm23.8-40.4c-.1-1.3-.1-2.3-.3-3.3-1.8-9.9-10.9-15.5-20.4-13.3-9.3 2.1-15.3 8-17.5 17.4-1.8 7.8 2 15.7 9.2 18.9 5.5 2.4 11 2.1 16.3-.6 7.9-4.1 12.2-10.5 12.7-19.1z"
                fill-rule="nonzero"
              />
            </g>
          </g>
        </svg>
        <div
          class="absolute inset-0 rounded-t-lg lg:rounded-t-none lg:rounded-r-lg shadow-[inset_0px_0px_0px_1px_rgba(26,26,0,0.16)] dark:shadow-[inset_0px_0px_0px_1px_#fffaed2d]"
        ></div>
      </div>
    </main>
  </div>
  <div class="hidden h-14.5 lg:block"></div>
</div>
