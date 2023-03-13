<script lang="ts">
  import Icon from "$lib/component/Icon.svelte";
import Github from "./Github.svelte";
  import Twitter from "./Twitter.svelte";

  let email: string = "";
  let password: string = "";
  let checkbox: boolean = false;
  let loading: boolean = false;
  function handleSubmit() {
    loading = true;
    console.log("Email:", email);
    console.log("Password:", password);
    console.log("Checkbox:", checkbox)
  }
</script>

<div class="w-full h-full flex justify-center items-center">
  <form
    on:submit|preventDefault={handleSubmit}
    autocomplete="off"
    class="w-full md:w-2/6 p-8 flex flex-col gap-4 justify-start rounded-lg md:border border-primary"
  >
  <div class="w-full inline-flex justify-start">
    <h2 class="text-3xl text-center font-bold mb-4">Log in</h2>
  </div>
    <div class="form-control w-full">
      <label class="label" for="email">
        <span class="label-text">What is your email?</span>
      </label>
      <input
        id="email"
        bind:value={email}
        name="email"
        type="email"
        placeholder="joe@email.com"
        class="input input-bordered w-full border-primary"
      />
    </div>
    <div class="form-control w-full">
      <label class="label" for="password">
        <span class="label-text">What is your password?</span>
      </label>
      <input
        id="password"
        bind:value={password}
        name="password"
        type="password"
        placeholder="*********"
        class="input input-bordered w-full border-primary"
      />
    </div>
    <div class="form-control">
      <label class="label cursor-pointer justify-start">
        <input type="checkbox" class="checkbox checkbox-primary mr-3" bind:checked={checkbox} />
        <span class="label-text select-none">Remember me</span>
      </label>
    </div>
    <div>
      <button type="submit" class:loading class="w-full btn btn-primary">Login</button>
    </div>
    <div class="divider select-none">Social Media</div>
    <div class="w-full flex justify-between gap-4 select-none">
      <div class="flex flex-row items-center gap-2 border border-primary-focus rounded p-3 hover:bg-primary-focus hover:text-primary-content">
        <div class="w-8 h-8 bg-slate-100 rounded-full inline-flex justify-center items-center p-1">
          <Twitter />
        </div>
        <div class="flex flex-col">
          <p class="font-mono text-xs">Sign in with</p>
          <span class="font-semibold text-lg">Twitter</span>
        </div>
      </div>
      <div class="flex flex-row items-center gap-2 border border-primary-focus rounded p-3 hover:bg-primary-focus hover:text-primary-content">
        <div class="w-8 h-8 bg-slate-100 rounded-full inline-flex justify-center items-center p-1">
          <Github />
        </div>
        <div class="flex flex-col">
          <p class="font-mono text-xs">Sign in with</p>
          <span class="font-semibold text-lg">Github</span>
        </div>
      </div>
      
    </div>
  </form>
</div>