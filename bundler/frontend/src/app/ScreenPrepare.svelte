<script lang="ts">
  export let progress: string[];
</script>

<div class="container">
  {#each progress as p}
    <p>{p}</p>
  {/each}
</div>

<style>
  /* .container {
  }

  .container > p {
  } */
</style>
