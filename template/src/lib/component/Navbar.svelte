<script>
    import ThemeSelect from "./ThemeSelect.svelte";
  
    import { page } from "$app/stores";
    import { fly } from "svelte/transition";
    import Icon from "./Icon.svelte";
  
    const routes = [
      {
        path: "/auth/magicway",
        icon: "home",
      },
      {
        path: "/auth/login",
        icon: "key",
      },
      {
        path: "/auth/register",
        icon: "person",
      },
    ];
  </script>
  <ThemeSelect />
  <div class="btm-nav md:left-nav">
    {#each routes as route}
      <a href={route.path} class:active={$page.url.pathname.startsWith(route.path)}>
        <Icon name={route.icon} type="filled" />
      </a>
    {/each}
  </div>
  <style global lang="postcss">
      .left-nav {
          @apply fixed z-40 right-0 w-12 h-screen flex flex-col justify-center items-center;
          @apply bg-primary;
      }
      .left-nav>*{
          @apply relative m-1 h-12;
      }
  </style>