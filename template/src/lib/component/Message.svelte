<script lang="ts">
    export let text: string
</script>

<div class="main">
    <h1>{text}</h1>
</div>

<style lang="postcss">
    .main{
        @apply h-screen w-screen flex justify-center items-center;
    }
    .main h1 {
        @apply font-semibold text-lg inline-flex justify-center items-center;
    }
</style>