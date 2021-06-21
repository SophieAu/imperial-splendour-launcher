<script lang="ts">
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
    <input bind:value={selectedFolder} /><button on:click={selectFolder}>Select Folder</button>
  </div>
  <p>New Version:</p>
  <div class="versionInput">
    <input bind:value={versionNumber} />
  </div>
  <p>File List:</p>
  <div class="fileSelect">
    <input bind:value={selectedFileListFile} /><button on:click={selectFile}>Select File</button>
  </div>
  <p>Package Raw Files?</p>
  <div>
    <input type="checkbox" bind:checked={shouldPackageRawFiles} />
  </div>
</div>

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
    display: grid;
    grid-template-columns: max-content max-content;
    grid-gap: 1.5rem 1rem;
    place-items: center left;
  }

  .inputGrid > p {
    margin: 0;
  }

</style>
