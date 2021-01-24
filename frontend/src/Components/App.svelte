<script>
  import Button from "./Button.svelte";
  import headerLogo from "../assets/hero_logo.png";

  let version;
  let pageTitle = "Imperial Splendour: Rise of the Republic";

  window.backend.version().then((result) => {
    version = result;
  });

  const handlePlay = () => {
    window.backend.play();
  };
  const handleSwitch = () => {
    window.backend.switchMode();
  };
  const handleWebsite = () => {
    window.backend.openWebsite();
  };
  const handleUninstall = () => {
    window.backend.uninstall();
  };
  const handleExit = () => {
    window.backend.ExitCall.Exit();
  };
</script>

<svelte:head>
  <title>{pageTitle}</title>
</svelte:head>
<main>
  <h1>
    <img src={headerLogo} title={pageTitle} alt={pageTitle} />
  </h1>
  <div class="buttonContainer">
    <Button text={"Play"} handleClick={handlePlay} />
    <Button text={"Switch"} handleClick={handleSwitch} />
    <Button text={"Website"} handleClick={handleWebsite} />
    <Button text={"Uninstall"} handleClick={handleUninstall} />
    <Button text={"Exit"} handleClick={handleExit} />
  </div>
  <footer><span class="prefix">v</span>{version}</footer>
</main>

<style>
  main {
    background: center / contain var(--img-bg);
    height: calc(100vw / 1.6);
    overflow: hidden;

    display: flex;
    flex-direction: column;
  }

  h1 {
    margin: 0 auto;
    padding: calc(var(--height-ratio) * 58) 0 calc(var(--height-ratio) * 128);
    display: grid;
    place-items: center;
  }

  h1 > img {
    width: calc(var(--width-ratio) * 1000);
  }

  .buttonContainer {
    flex: 1;
    display: flex;
    justify-content: space-evenly;
    align-items: flex-start;
  }

  footer {
    text-align: right;
    margin-bottom: -0.75rem;
    padding-right: 0.25rem;
    font-family: "IM FELL English";
    font-size: 4rem;
  }

  footer > .prefix {
    font-size: 2rem;
  }
</style>
