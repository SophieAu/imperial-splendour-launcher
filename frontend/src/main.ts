// @ts-check
import Wails from '@wailsapp/runtime';

import Root from './Root.svelte';

let app;

Wails.Init(() => {
  app = new Root({ target: document.body });
});

export default app;
