<script context="module" lang="ts">
  export type APIType = {
    Version: () => Promise<string>;
    IsActive: () => Promise<boolean>;
    Play: () => Promise<void>;
    Switch: () => Promise<void>;
    GoToWebsite: () => Promise<void>;
    Uninstall: () => Promise<void>;
    Exit: () => Promise<void>;
  };
</script>

<script lang="ts">
  import { onMount } from 'svelte';

  import Button from './Button.svelte';
  import isLogo from './assets/hero_logo_is.png';
  import etwLogo from './assets/hero_logo_etw.png';
  import Modal from './Modal.svelte';
  import { pageTitle, etwTitle, apiErrors } from './strings';

  let version = '';
  let isISActive = true;
  let errorMessage = '';

  export let API: APIType;

  const callAPI = async (callback: () => Promise<void>, errorCode: keyof typeof apiErrors) => {
    try {
      await callback();
    } catch (e) {
      console.log(e);
      errorMessage = apiErrors[errorCode];
    }
  };

  onMount(async () => {
    await callAPI(async () => {
      version = await API.Version();
      isISActive = await API.IsActive();
    }, 'startup');
  });

  const switchError = () => (isISActive ? 'switchToETW' : 'switchToIS');
  const switchMode = async () => {
    await API.Switch();
    isISActive = await API.IsActive();
  };

  const dismissError = () => {
    errorMessage = '';
  };
</script>

<svelte:head>
  <title>{pageTitle}</title>
</svelte:head>
<main>
  <h1>
    <img src={isISActive ? isLogo : etwLogo} alt={isISActive ? pageTitle : etwTitle} />
  </h1>
  <div class="buttonContainer">
    <Button text={'Play'} onClick={async () => callAPI(API.Play, 'play')} />
    <Button text={'Switch'} onClick={async () => callAPI(switchMode, switchError())} />
    <Button text={'Website'} onClick={async () => callAPI(API.GoToWebsite, 'website')} />
    <Button text={'Uninstall'} onClick={async () => callAPI(API.Uninstall, 'uninstall')} />
    <Button text={'Exit'} onClick={async () => callAPI(API.Exit, 'exit')} />
  </div>
  <footer>
    <span class="prefix">v</span><span class="version">{version}</span>
  </footer>
  {#if errorMessage}
    <Modal message={errorMessage} onClick={dismissError} />
  {/if}
</main>

<style>
  main {
    background: center / contain var(--img-bg);
    height: calc(100vw / 1.6);
    overflow: hidden;

    display: flex;
    flex-direction: column;
  }

  h1 {
    margin: 0 auto;
    padding: calc(var(--height-ratio) * 58) 0 calc(var(--height-ratio) * 128);
    display: grid;
    place-items: center;
  }

  h1 > img {
    width: calc(var(--width-ratio) * 1000);
    height: calc(var(--width-ratio) * 374);
    object-fit: contain;
  }

  .buttonContainer {
    flex: 1;
    display: flex;
    justify-content: space-evenly;
    align-items: flex-start;
  }

  footer {
    text-align: right;
    margin-bottom: -0.25rem;
    padding-right: 0.25rem;
    font-family: var(--font-family-body);
  }

  footer > .prefix {
    font-size: var(--font-size-body);
  }

  footer > .version {
    font-size: var(--font-size);
  }
</style>
