<script>
  export let text;
  export let onClick;
</script>

<button on:click={onClick}>{text}</button>

<style>
  button {
    --download-button-height: calc(var(--height-ratio) * 100);
    --download-button-width: calc(var(--width-ratio) * 290);
    --color-link-hover: #5c4e3c;

    background: no-repeat center/100% var(--button-texture);
    font: normal var(--font-size) / var(--download-button-height) var(--font-family-heading);

    text-align: center;
    text-decoration: none;
    width: var(--download-button-width);
    padding: 0;
    margin: 0;
    border: none;
  }

  button:hover {
    color: var(--color-link-hover);
    cursor: pointer;
  }
</style>
