<script>
  import Button from "./Button.svelte";

  let version;
  let pageTitle = "Imperial Splendour: Rise of the Republic";

  window.backend.version().then((result) => {
    version = result;
  });

  const handleExit = () => {
    window.backend.ExitCall.Exit();
  };

  const handlePlay = () => {};
  const handleSwitch = () => {};
  const handleWebsite = () => {};
  const handleUninstall = () => {};
</script>

<svelte:head>
  <title>{pageTitle}</title>
</svelte:head>
<main>
  <h1><span>{pageTitle}</span></h1>
  <div class="buttonContainer">
    <Button text={"Play"} handleClick={handlePlay} />
    <Button text={"Switch"} handleClick={handleSwitch} />
    <Button text={"Website"} handleClick={handleWebsite} />
    <Button text={"Uninstall"} handleClick={handleUninstall} />
    <Button text={"Exit"} handleClick={handleExit} />
  </div>
  <footer>v{version}</footer>
</main>

<style>
  main {
    background: center / contain url("/assets/background.png");
    padding: 1em;
    min-height: calc((100vw / 1.6) - 2em);
  }

  h1 {
    background: center / contain no-repeat url("/assets/hero_logo.png");

    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
    height: 200px;
  }

  h1 > span {
    display: none;
  }

  .buttonContainer {
    display: flex;
  }

  footer {
    position: absolute;
    bottom: 0;
    right: 0;
  }
</style>
