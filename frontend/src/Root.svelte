<script lang="ts">
  import { injectGlobal } from '@emotion/css/dist/emotion-css.umd.min.js';

  import imFellReg2 from './assets/imfellenglish.woff2';
  import imFellSC2 from './assets/imfellenglishsc.woff2';

  import App from './App.svelte';

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

  type APIType = {
    Version: () => Promise<string>;
    IsActive: () => Promise<boolean>;
    Play: () => Promise<void>;
    Switch: () => Promise<void>;
    GoToWebsite: () => Promise<void>;
    Uninstall: () => Promise<void>;
    Exit: () => Promise<void>;
  };

  const API = (window as any)?.backend?.API as APIType;
</script>

<div>
  <App {API} />
</div>
