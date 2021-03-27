<script lang="ts">
  import { onMount } from 'svelte';
  import { css, injectGlobal } from '@emotion/css/dist/emotion-css.umd.min.js';

  import imFellReg2 from './assets/imfellenglish.woff2';
  import imFellSC2 from './assets/imfellenglishsc.woff2';
  import background from './assets/background.png';
  import textureBtn from './assets/texture_btn.png';
  import textureBg from './assets/texture_bg.jpg';

  import App from './App.svelte';

  type APIType = {
    Version: () => Promise<string>;
    IsActive: () => Promise<boolean>;
    Play: () => Promise<void>;
    Switch: () => Promise<void>;
    GoToWebsite: () => Promise<void>;
    Uninstall: () => Promise<void>;
    Exit: () => Promise<void>;
  };

  const rootStyle = css`
    --img-bg: url(${background});
    --button-texture: url(${textureBtn});
    --modal-bg: url(${textureBg});
  `;

  const API = (window as any)?.backend?.API as APIType;

  injectGlobal`
    @font-face {
      font-family: 'IM FELL English';
      src: url(${imFellReg2}) format('woff2');
      font-weight: 'normal';
      font-style: 'normal';
    }

  @font-face {
    font-family: 'IM FELL English SC';
    src: url(${imFellSC2}) format('woff2');
    font-weight: 'normal';
    font-style: 'normal';
  }
  `;
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
