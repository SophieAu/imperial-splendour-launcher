<script lang="ts">
  import { mapError } from './helpers';
  import InputForm from './InputForm.svelte';

  import Modal from './Modal.svelte';
  import { pageTitle } from './strings';
  import type { APIType } from './types';

  export let API: APIType;

  let modalText = '';

  let selectedFolder = '';
  let selectedFileListFile = '';
  let versionNumber = '';
  let shouldPackageRawFiles = false;

  const prepareBundling = async () => {
    try {
      await API.Prepare(selectedFolder, versionNumber, shouldPackageRawFiles, selectedFileListFile);
    } catch (e: unknown) {
      modalText = mapError(e as Error);
    }
  };

  const bundle = async () => {
    try {
      await API.Bundle();
    } catch (e: unknown) {
      modalText = mapError(e as Error);
    }
  };

  const dismissModal = () => {
    modalText = '';
  };

  const selectFolder = async () => {
    try {
      selectedFolder = await API.SelectSourceDir();
    } catch (e: unknown) {
      modalText = mapError(e as Error);
    }
  };

  const selectFile = async () => {
    try {
      selectedFolder = await API.SelectFileListLocation();
    } catch (e: unknown) {
      modalText = mapError(e as Error);
    }
  };

</script>

<svelte:head>
  <title>{pageTitle}</title>
</svelte:head>
<main class="root">
  <h1 class="heading">Imperial Splendour Bundler</h1>
  <InputForm
    bind:selectedFolder
    bind:selectedFileListFile
    bind:versionNumber
    bind:shouldPackageRawFiles
    {selectFile}
    {selectFolder}
  />
  <div class="buttonContainer">
    <button on:click={prepareBundling}>Bundle</button>
  </div>
  {#if modalText}
    <Modal bind:message={modalText} onClick={dismissModal} />
  {/if}
</main>

<style>
  .root {
    width: 100vw;
    height: 100vh;
    overflow: hidden;

    display: flex;
    flex-direction: column;
    justify-content: space-between;

    box-sizing: border-box;
    padding: 1rem;
  }

  .heading {
    margin: 0;
    font-size: 1.5rem;
  }

  .buttonContainer {
    align-self: flex-end;
  }

</style>
