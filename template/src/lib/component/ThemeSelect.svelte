<script lang="ts">
  import { theme } from "$lib/store/client";
  import { set } from "$lib/cookie";
  import Icon from "./Icon.svelte";
  const t = [
    "light",
    "dark",
    "cupcake",
    "bumblebee",
    "emerald",
    "corporate",
    "synthwave",
    "retro",
    "cyberpunk",
    "valentine",
    "halloween",
    "garden",
    "forest",
    "aqua",
    "lofi",
    "pastel",
    "fantasy",
    "wireframe",
    "black",
    "luxury",
    "dracula",
    "cmyk",
    "autumn",
    "business",
    "acid",
    "lemonade",
    "night",
    "coffee",
    "winter",
  ];
  const h = (v: string) => {
    $theme = v;
    document.documentElement.setAttribute("data-theme", $theme);
    set("theme", v);
  };
</script>

<!-- The button to open modal -->
<label for="my-modal-4" class="btn btn-square">
  <Icon name="layers" type="filled" class="text-xl" />
</label>

<!-- Put this part before </body> tag -->
<input type="checkbox" id="my-modal-4" class="modal-toggle" />
<label for="my-modal-4" class="modal cursor-pointer">
  <label class="modal-box relative rounded-tr-none rounded-br-none" for="">
    <div class="grid grid-cols-2 gap-3 p-3" data-focustheme="true">
      {#each t as i}
        <button
          type="button"
          on:click={() => h(i)}
          class="outline-base-content overflow-hidden rounded-lg text-left"
          data-set-theme={i}
        >
          <div
            data-theme={i}
            class="bg-base-100 text-base-content w-full cursor-pointer font-sans"
          >
            <div class="grid grid-cols-5 grid-rows-3">
              <div
                class="col-span-5 row-span-3 row-start-1 flex gap-2 py-3 px-4 items-center"
              >
                {#if $theme === i}<Icon
                    name="done"
                    type="filled"
                    class="text-sm"
                  />{/if}
                <div class="flex-grow text-sm font-bold">{i}</div>
                <div class="flex flex-shrink-0 flex-wrap gap-1 h-full">
                  <div class="bg-primary w-2 rounded" />
                  <div class="bg-secondary w-2 rounded" />
                  <div class="bg-accent w-2 rounded" />
                  <div class="bg-neutral w-2 rounded" />
                </div>
              </div>
            </div>
          </div>
        </button>
      {/each}
    </div>
  </label>
</label>
