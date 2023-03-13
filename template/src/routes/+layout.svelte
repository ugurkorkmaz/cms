<script lang="ts">
  import "../app.postcss";
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Body from "$lib/component/Body.svelte";
  import Footer from "$lib/component/Footer.svelte";
  import Header from "$lib/component/Header.svelte";
  import Transition from "$lib/component/Transition.svelte";
  import Wrapper from "$lib/component/Wrapper.svelte";
  import { theme, routes, type Route } from "$lib/store/client";
  import { get } from "$lib/cookie";
  import Navbar from "$lib/component/Navbar.svelte";
  import { guard } from "$lib/store/auth";

  $: if (browser) {
    const current = $routes.find((route) => route.path === $page.url.pathname) as Route
    if ($guard && current.auth) goto(current.path);
    $theme = get("theme") as string;
  }
</script>
<Wrapper>
  <Header />
  <Navbar />
  <Transition url={$page.url}>
    <Body>
      <slot />
    </Body>
  </Transition>
  <Footer />
</Wrapper>
