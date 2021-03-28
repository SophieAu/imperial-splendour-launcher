<script lang="ts">
  import { onMount } from 'svelte';

  import Button from './Button.svelte';
  import isLogo from './assets/hero_logo_is.png';
  import etwLogo from './assets/hero_logo_etw.png';
  import Modal from './Modal.svelte';
  import { pageTitle, etwTitle, apiErrors, versionPrefix } from './strings';
  import * as styles from './styles.app';

  type APIType = {
    Version: () => Promise<string>;
    IsActive: () => Promise<boolean>;
    Play: () => Promise<void>;
    Switch: () => Promise<void>;
    GoToWebsite: () => Promise<void>;
    Uninstall: () => Promise<void>;
    Exit: () => Promise<void>;
  };

  let version = '';
  let isISActive = true;
  let errorMessage = '';

  export let API: APIType;

  const callAPI = async (callback: () => Promise<void>, errorCode: keyof typeof apiErrors) => {
    try {
      await callback();
    } catch (e) {
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
<main class={styles.root}>
  {#if isISActive !== undefined}
    <h1 class={styles.heading}>
      <img src={isISActive ? isLogo : etwLogo} alt={isISActive ? pageTitle : etwTitle} />
    </h1>
    <div class={styles.buttonContainer}>
      <Button text={'Play'} onClick={async () => callAPI(API.Play, 'play')} />
      <Button text={'Switch'} onClick={async () => callAPI(switchMode, switchError())} />
      <Button text={'Website'} onClick={async () => callAPI(API.GoToWebsite, 'website')} />
      <Button text={'Uninstall'} onClick={async () => callAPI(API.Uninstall, 'uninstall')} />
      <Button text={'Exit'} onClick={async () => callAPI(API.Exit, 'exit')} />
    </div>
    <footer class={styles.footer}>
      <span class="prefix">{versionPrefix}</span><span class="version">{version}</span>
    </footer>
  {/if}
  {#if errorMessage}
    <Modal message={errorMessage} onClick={dismissError} />
  {/if}
</main>
