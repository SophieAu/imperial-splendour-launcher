<script lang="ts">
  import { injectGlobal } from '@emotion/css/dist/emotion-css.umd.min.js';

  import App from './App.svelte';
  import imFellReg from './assets/imfellenglish.woff';
  import imFellReg2 from './assets/imfellenglish.woff2';
  import imFellSC from './assets/imfellenglishsc.woff';
  import imFellSC2 from './assets/imfellenglishsc.woff2';
  import type { APIType } from './types';

  injectGlobal`
  :root {
    font-size: 16px;
  }

  body {
    overflow: hidden;
    margin: 0;
    height: 100vh;
    width: 100vw;
  }

  @font-face {
    font-family: 'IM FELL English';
    src: url(${imFellReg2}) format('woff2'), url(${imFellReg}) format('woff');
    font-weight: 'normal';
    font-style: 'normal';
  }

  @font-face {
    font-family: 'IM FELL English SC';
    src: url(${imFellSC2}) format('woff2'), url(${imFellSC}) format('woff');
    font-weight: 'normal';
    font-style: 'normal';
  }
  `;

  const API = (window as any)?.backend?.API as APIType;
</script>

<div>
  <App {API} />
</div>
