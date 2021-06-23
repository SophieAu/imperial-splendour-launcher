<script lang="ts">
  import App from './App.svelte';
  import type { APIType } from '../types';
  import runtime from '@wailsapp/runtime';

  const API = (window as any)?.backend?.API as APIType;
  const Store = runtime.Store.New('Log');
</script>

<div>
  <App {API} {Store} />
</div>

<style>
  :global(:root) {
    font-size: 16px;
    font-family: system;
  }

  :global(body) {
    overflow: hidden;
    margin: 0;
    height: 26.25rem;
    width: 33.75rem;
  }
</style>
