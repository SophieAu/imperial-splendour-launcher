<script lang="ts">
  import { mapError } from '../helpers';
  import InputForm from './ScreenInput.svelte';
  import BundleProgress from './ScreenBundle.svelte';
  import ErrorScreen from './ScreenError.svelte';
  import PrepProgress from './ScreenPrepare.svelte';
  import runtime from '@wailsapp/runtime';

  import { pageTitle } from '../strings';
  import type { APIType } from '../types';
  import { onMount } from 'svelte';

  enum STAGE {
    INPUT,
    PREPARE,
    BUNDLE,
    ERROR,
  }

  export let API: APIType;

  const versionRegex = new RegExp(/^\d+.\d+(.\d+)?$/);

  let modalText = '';
  let errorMessage = '';
  let stage = STAGE.INPUT;
  let isButtonEnabled = false;

  let selectedFolder = '';
  let selectedFileListFile = '';
  let versionNumber = '';
  let shouldPackageRawFiles = false;

  let compileProgressInfo = '';

  let progressInfo: string[] = [];
  onMount(() => {
    runtime.Store.New('Log').subscribe((newProgressInfo: string[]) => {
      progressInfo = newProgressInfo;
    });
  });

  const prepareBundling = async () => {
    let modalMsg: string[] = [];
    if (!selectedFolder) {
      modalMsg.push('You need to select a source folder directory.');
    }
    if (!selectedFileListFile) {
      modalMsg.push('You need to select a file list file.');
    }
    if (!versionRegex.test(versionNumber)) {
      modalMsg.push('The version needs to be of the format x.y(.z).');
    }
    if (!selectedFolder || !selectedFileListFile || !versionRegex.test(versionNumber)) {
      modalText = modalMsg.join(' ');
      return;
    }

    stage = STAGE.PREPARE;
    isButtonEnabled = false;

    try {
      await API.Prepare(selectedFolder, versionNumber, shouldPackageRawFiles, selectedFileListFile);
      isButtonEnabled = true;
    } catch (e: unknown) {
      errorMessage = mapError(e as Error);
      stage = STAGE.ERROR;
    }
  };

  const bundle = async () => {
    stage = STAGE.BUNDLE;
    isButtonEnabled = false;
    try {
      compileProgressInfo = 'Compiling. Check the log for more info';
      await API.Bundle();
      compileProgressInfo = "Finished Compiling. You're all ready to go.";
      isButtonEnabled = true;
    } catch (e: unknown) {
      errorMessage = mapError(e as Error);
      stage = STAGE.ERROR;
    }
  };

  const exit = async () => {
    try {
      await API.Exit();
    } catch (e: unknown) {
      modalText = mapError(e as Error);
    }
  };

  const dismissModal = async () => {
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
  {#if stage === STAGE.INPUT}
    <InputForm
      bind:selectedFolder
      bind:selectedFileListFile
      bind:versionNumber
      bind:shouldPackageRawFiles
      {selectFile}
      {selectFolder}
      bind:modalText
      {dismissModal}
    />
    <div class="buttonContainer">
      <button on:click={prepareBundling}>Bundle</button>
    </div>
  {/if}
  {#if stage === STAGE.PREPARE}
    <PrepProgress bind:progress={progressInfo} />
    <div class="buttonContainer">
      <button on:click={prepareBundling} disabled={!isButtonEnabled}>Compile</button>
    </div>
  {/if}
  {#if stage === STAGE.BUNDLE}
    <BundleProgress bind:progress={compileProgressInfo} />
    <div class="buttonContainer">
      <button on:click={bundle} disabled={!isButtonEnabled}>Finish</button>
    </div>
  {/if}
  {#if stage === STAGE.ERROR}
    <ErrorScreen bind:message={errorMessage} />
    <div class="buttonContainer">
      <button on:click={exit}>Close</button>
    </div>
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
