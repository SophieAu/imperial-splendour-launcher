<script>
  import App from './App.svelte';
  import { root } from './rootStyles.js';
  import imFellReg2 from './assets/imfellenglish.woff2';
  import imFellSC2 from './assets/imfellenglishsc.woff2';
  import { onMount } from 'svelte';

  const { API } = window.backend || { API: undefined };

  onMount(async () => {
    const src = (url) => `url(${url}) format('woff2')`;
    const descriptors = { weight: 'normal', style: 'normal' };

    const reg = new FontFace('IM FELL English', src(imFellReg2), descriptors);
    const sc = new FontFace('IM FELL English SC', src(imFellSC2), descriptors);
    await reg.load();
    await sc.load();

    document.fonts.add(reg);
    document.fonts.add(sc);
  });
</script>

<div class={root}>
  <App {API} />
</div>

<style>
  :global(body) {
    --width-ratio: calc(100vw / 1920);
    --height-ratio: calc(100vh / 1200);

    overflow: hidden;
    margin: 0;
    height: 100vh;
    width: 100vw;
  }
</style>
