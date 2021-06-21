<script lang="ts">
  import { mapError } from './helpers';

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
  <div class="inputGrid">
    <p>Source File Path:</p>
    <div>
      <input bind:value={selectedFolder} /><button on:click={selectFolder}>Select Folder</button>
    </div>
    <p>New Version:</p>
    <div>
      <input bind:value={versionNumber} />
    </div>
    <p>File List:</p>
    <div>
      <input /><button on:click={selectFile}>Select File</button>
    </div>
    <p>Package Raw Files?</p>
    <div>
      <input type="checkbox" bind:checked={shouldPackageRawFiles} />
    </div>
  </div>
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
  }

  .heading {
  }

  .inputGrid {
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .buttonContainer {
  }

</style>
