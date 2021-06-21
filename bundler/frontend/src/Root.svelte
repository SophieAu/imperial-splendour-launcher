<script lang="ts">
  import App from './App.svelte';
  import type { APIType } from './types';
  import runtime from '@wailsapp/runtime';
  import { onMount } from 'svelte';

  const API = (window as any)?.backend?.API as APIType;
  let stateStore;
  let progressInfo: string[] = [];

  onMount(() => {
    stateStore = runtime.Store.New('State');

    stateStore.subscribe((newProgressInfo: string[]) => {
      progressInfo = newProgressInfo;
    });
  });

</script>

<div>
  <App {API} />
</div>

<style>
  :global(:root) {
    font-size: 16px;
  }

  :global(body) {
    overflow: hidden;
    margin: 0;
    height: 100vh;
    width: 100vw;
  }

</style>
