<script>
  export let text;
  export let handleClick;
</script>

<button on:click={handleClick}>{text}</button>

<style>
  button {
    --font-family-heading: "IM FELL English SC";
    --download-button-height: 6.25rem;
    --download-button-height: calc(var(--height-ratio) * 100);
    --download-button-width: 18.125rem;
    --color-link-hover: #5c4e3c;

    background: no-repeat center/100% var(--button-texture);
    font: normal 2.5rem / var(--download-button-height)
      var(--font-family-heading);
    height: var(--download-button-height);
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
