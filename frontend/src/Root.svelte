<script lang="ts">
  import { onMount } from 'svelte';
  import { css } from '@emotion/css/dist/emotion-css.umd.min.js';

  import imFellReg2 from './assets/imfellenglish.woff2';
  import imFellSC2 from './assets/imfellenglishsc.woff2';
  import background from './assets/background.png';
  import textureBtn from './assets/texture_btn.png';
  import textureBg from './assets/texture_bg.jpg';

  import App, { APIType } from './App.svelte';

  const rootStyle = css`
    --img-bg: url(${background});
    --button-texture: url(${textureBtn});
    --modal-bg: url(${textureBg});
  `;

  const API = (window as any)?.backend?.API as APIType;

  onMount(async () => {
    const src = (url: string) => `url(${url}) format('woff2')`;
    const descriptors = { weight: 'normal', style: 'normal' };

    const reg = new FontFace('IM FELL English', src(imFellReg2), descriptors);
    const sc = new FontFace('IM FELL English SC', src(imFellSC2), descriptors);
    await reg.load();
    await sc.load();

    document.fonts.add(reg);
    document.fonts.add(sc);
  });
</script>

<div class={rootStyle}>
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

    --font-size-factor: calc(40 * calc((100vh - 800px) / (1200 - 800)));
    --font-size: clamp(32px, var(--font-size-factor), 40px);
    --font-family-heading: 'IM FELL English SC';

    --font-size-factor-body: calc(28 * calc((100vh - 800px) / (1200 - 800)));
    --font-size-body: clamp(20px, var(--font-size-factor), 28px);
    --font-family-body: 'IM FELL English';
  }
</style>
