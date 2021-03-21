import Root from './Root.svelte';

import Wails from '@wailsapp/runtime';

let app: Root | undefined;

Wails.Init(() => {
  app = new Root({ target: document.body });
});

export default app;
