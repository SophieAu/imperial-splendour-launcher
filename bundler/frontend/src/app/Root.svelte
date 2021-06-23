<script lang="ts">
  import App from './App.svelte';
  import type { APIType } from '../types';

  const API = (window as any)?.backend?.API as APIType;
</script>

<div>
  <App {API} />
</div>

<style>
  :global(:root) {
    font-size: 16px;
    font-family: system;
  }

  :global(body) {
    overflow: hidden;
    margin: 0;
    height: 100vh;
    width: 100vw;
  }
</style>
