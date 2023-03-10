<script lang="ts">
  import Message from "$lib/component/Message.svelte";
  import type { PageData } from "./$types";

  export let data: PageData;
  console.log(data)
</script>
 
<Message text={data?.meta.title} />
