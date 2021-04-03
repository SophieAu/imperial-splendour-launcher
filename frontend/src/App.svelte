<script lang="ts">
  import { onMount } from 'svelte';

  import etwLogo from './assets/hero_logo_etw.png';
  import isLogo from './assets/hero_logo_is.png';
  import Button from './Button.svelte';
  import { getNewestVersion } from './helpers';
  import Modal from './Modal.svelte';
  import { apiErrors, etwTitle, newVersion, pageTitle, versionPrefix } from './strings';
  import * as styles from './styles.app';
  import type { APIType } from './types';

  let version = '';
  let isISActive = true;
  let errorMessage = '';

  export let API: APIType;

  const callAPI = async (callback: () => Promise<void>, errorCode: keyof typeof apiErrors) => {
    try {
      await callback();
    } catch (e) {
      e.message();
      errorMessage = apiErrors[errorCode];
    }
  };

  onMount(async () => {
    await callAPI(async () => {
      version = await API.Version();
      isISActive = await API.IsActive();
    }, 'startup');

    try {
      const newestVersion = await getNewestVersion();

      if (version < newestVersion) {
        errorMessage = newVersion;
      }
    } catch {}
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
