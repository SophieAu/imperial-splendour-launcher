import Root from './Root.svelte';

import * as Wails from '@wailsapp/runtime';

let app;

Wails.Init(() => {
  app = new Root({ target: document.body });
});

export default app;
