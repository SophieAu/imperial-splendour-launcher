<script lang="ts">
  import Modal from './Modal.svelte';

  export let modalText = '';
  export let dismissModal: () => void;

  export let selectedFolder: string;
  export let selectedFileListFile: string;
  export let versionNumber: string;
  export let shouldPackageRawFiles: boolean;

  export let selectFolder: () => void;
  export let selectFile: () => void;
</script>

<div class="inputGrid">
  <p>Source File Path:</p>
  <div class="fileSelect">
    <input bind:value={selectedFolder} required /><button on:click={selectFolder}
      >Select Folder</button>
  </div>
  <p>New Version:</p>
  <div class="versionInput">
    <input bind:value={versionNumber} required pattern="\d+.\d+(.\d+)?" placeholder="e.g. 2.9" />
  </div>
  <p>File List:</p>
  <div class="fileSelect">
    <input bind:value={selectedFileListFile} required /><button on:click={selectFile}
      >Select File</button>
  </div>
  <p>Package Raw Files?</p>
  <div>
    <input type="checkbox" bind:checked={shouldPackageRawFiles} />
  </div>
</div>
{#if modalText}
  <Modal bind:message={modalText} onClick={dismissModal} />
{/if}

<style>
  .fileSelect > input {
    width: 16rem;
  }
  .fileSelect > button {
    width: 6rem;
  }

  .versionInput > input {
    width: 5rem;
  }

  .inputGrid {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-items: left;
  }
  .inputGrid > p {
    margin: 0.75rem 1rem 0.75rem 0;
    width: 8.25rem;
  }
  .inputGrid > div {
    width: 22.5rem;
  }
</style>
