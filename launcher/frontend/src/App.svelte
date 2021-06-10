<script lang="ts">
  import { onMount } from 'svelte';

  import etwLogo from './assets/hero_logo_etw.png';
  import isLogo from './assets/hero_logo_is.png';
  import Button from './Button.svelte';
  import type { EndpointKeys } from './helpers';
  import { callAPI, getNewestVersion, mapSwitchError, mapUninstallError } from './helpers';
  import Modal from './Modal.svelte';
  import { apiErrors, etwTitle, newVersionAvailable, pageTitle, versionPrefix } from './strings';
  import * as styles from './styles.app';
  import type { APIType } from './types';

  export let API: APIType;

  let version = '';
  let isISActive: boolean | undefined = undefined;
  let modalText = '';

  const callEndpoint = callAPI(API);

  onMount(async () => {
    try {
      version = await API.Version();
      isISActive = await API.IsActive();

      try {
        const newestVersion = await getNewestVersion();
        if (!!newestVersion && version != newestVersion) modalText = newVersionAvailable;
      } catch {}
    } catch (e: unknown) {
      modalText = apiErrors.startup;
    }
  });

  const callSwitch = async () => {
    const isDeactivating = isISActive;
    let newError = '';

    try {
      await API.Switch();
    } catch (e: unknown) {
      newError = mapSwitchError(!!isDeactivating, e as Error);
    } finally {
      try {
        isISActive = await API.IsActive();
      } catch {
        newError = apiErrors.unexpectedOnSwitch;
      }
    }
    modalText = newError;
  };

  const callUninstall = async () => {
    try {
      await API.Uninstall();
    } catch (e: unknown) {
      modalText = mapUninstallError(e as Error);
    }
  };

  const call = async (key: EndpointKeys) => {
    try {
      await callEndpoint(key);
    } catch (e) {
      modalText = (e as Error).message;
    }
  };

  const dismissModal = () => {
    modalText = '';
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
      <Button text={`Play ${isISActive ? 'Rotr' : 'etw'}`} onClick={async () => call('play')} />
      <Button text={'Switch'} onClick={async () => callSwitch()} />
      <Button text={'Website'} onClick={async () => call('website')} />
      <Button text={'Uninstall'} onClick={async () => callUninstall()} />
      <Button text={'Exit'} onClick={async () => call('exit')} />
    </div>
    <footer class={styles.footer}>
      <span class="prefix">{versionPrefix}</span><span class="version">{version}</span>
    </footer>
  {/if}
  {#if modalText}
    <Modal bind:message={modalText} onClick={dismissModal} />
  {/if}
</main>
