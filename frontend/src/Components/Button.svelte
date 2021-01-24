<script>
  export let text;
  export let handleClick;
</script>

<button on:click={handleClick}>{text}</button>

<style>
  button {
    --font-family-heading: "IM FELL English SC";
    --download-button-height: calc(var(--height-ratio) * 100);
    --download-button-width: calc(var(--width-ratio) * 290);
    --color-link-hover: #5c4e3c;
    --font-size: calc(40 * calc((100vh - 800px) / (1200 - 800)));

    background: no-repeat center/100% var(--button-texture);
    font: normal clamp(32px, var(--font-size), 40px) /
      var(--download-button-height) var(--font-family-heading);

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
