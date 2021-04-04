<script lang="ts">
  import { onMount } from 'svelte';

  import etwLogo from './assets/hero_logo_etw.png';
  import isLogo from './assets/hero_logo_is.png';
  import Button from './Button.svelte';
  import { callAPI, EndpointKeys, getNewestVersion } from './helpers';
  import Modal from './Modal.svelte';
  import {
    apiErrorCodes,
    apiErrors,
    etwTitle,
    newVersion,
    pageTitle,
    versionPrefix,
  } from './strings';
  import * as styles from './styles.app';
  import type { APIType } from './types';

  export let API: APIType;

  let version = '';
  let isISActive: boolean | undefined = undefined;
  let errorMessage = '';

  const callEndpoint = callAPI(API);

  onMount(async () => {
    try {
      version = await API.Version();
      isISActive = await API.IsActive();

      try {
        const newestVersion = await getNewestVersion();
        if (newestVersion && version != newestVersion) errorMessage = newVersion(newestVersion);
      } catch {}
    } catch (e: unknown) {
      errorMessage = apiErrors['startup'];
    }
  });

  const callSwitch = async () => {
    try {
      await API.Switch();
      isISActive = await API.IsActive();
    } catch (e: unknown) {
      // (e as Error).message;
      errorMessage = apiErrors[isISActive ? 'switchToETW' : 'switchToIS'];
    }
  };

  const callUninstall = async () => {
    try {
      await API.Uninstall();
      // isISActive = await API.IsActive();
    } catch (e: unknown) {
      // (e as Error).message;
      errorMessage = apiErrors['uninstall'];
    }
  };

  const call = async (key: EndpointKeys) => {
    try {
      await callEndpoint(key);
    } catch (e) {
      errorMessage = (e as Error).message;
    }
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
      <Button text={'Play'} onClick={async () => call('play')} />
      <Button text={'Switch'} onClick={async () => callSwitch()} />
      <Button text={'Website'} onClick={async () => call('website')} />
      <Button text={'Uninstall'} onClick={async () => callUninstall()} />
      <Button text={'Exit'} onClick={async () => call('exit')} />
    </div>
    <footer class={styles.footer}>
      <span class="prefix">{versionPrefix}</span><span class="version">{version}</span>
    </footer>
  {/if}
  {#if errorMessage}
    <Modal bind:message={errorMessage} onClick={dismissError} />
  {/if}
</main>
