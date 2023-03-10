<script lang="ts">
  import Icon from "$lib/component/Icon.svelte";
  import type { Snapshot } from "../../routes/$types";
  import Code from "./Code.svelte";
  import Image from "./Image.svelte";
  import Link from "./Link.svelte";
  import Table from "./Table.svelte";
  import type { text } from "./type";

  let papers: Array<text> = [
    {
      html: "<p>Hello World <p/>",
      markdown: "",
    },
  ];
  export const snapshot: Snapshot = {
    capture: () => papers,
    restore: (value: text[]) => (papers = value),
  };
</script>

<div class="editor">
  <div class="preview">
    {#each papers as item}
      {@html item.html}
    {/each}
  </div>
  <div class="tools">
    <Image />
    <Code />
    <Link />
    <Table />
  </div>
  <div
    class="writter"
    contenteditable="true"
    placeholder="Yeni paragrafı giriniz"
  />
  <div class="new-block">
    <button class="btn btn-circle">
      <Icon name="add_box" type="filled" />
    </button>
  </div>
</div>

<style global lang="postcss">
  .editor {
    @apply w-full h-full flex flex-col;
  }
  .editor .preview {
    @apply w-full h-auto;
  }
  .editor .writerr {
    @apply w-full h-96 border p-4 mb-2;
  }
  .editor .tools {
    @apply w-full h-8 flex justify-end md:justify-start gap-3;
  }
  .editor .new-block {
    @apply w-full h-8 flex justify-end md:justify-start;
  }
</style>
