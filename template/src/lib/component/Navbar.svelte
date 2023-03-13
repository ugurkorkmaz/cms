<script>
  import ThemeSelect from "./ThemeSelect.svelte";
  import { page } from "$app/stores";
  import Icon from "./Icon.svelte";
  import { routes } from "$lib/store/client";
</script>

<div class="fixed top-0 right-0">
  <ThemeSelect />
</div>
