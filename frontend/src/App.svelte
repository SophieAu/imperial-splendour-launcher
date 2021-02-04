<script>
  import Button from "./Button.svelte";
  import isLogo from "./assets/hero_logo_is.png";
  import etwLogo from "./assets/hero_logo_etw.png";

  let pageTitle = "Imperial Splendour: Rise of the Republic";
  let version = "";
  let isISActive = undefined;

  const { API } = window.backend;

  API.Version().then((result) => {
    version = result;
  });

  API.IsActive().then((result) => {
    isISActive = result;
  });

  const handlePlay = () => {
    API.Play();
  };
  const handleSwitch = async () => {
    try {
      await API.Switch();
      console.log("switched");
    } catch {
      console.log("oh noes");
    }
    API.IsActive().then((result) => {
      isISActive = result;
    });
  };
  const handleWebsite = () => {
    API.GoToWebsite();
  };
  const handleUninstall = () => {
    API.Uninstall();
  };
  const handleExit = () => {
    API.Exit();
  };
</script>

<svelte:head>
  <title>{pageTitle}</title>
</svelte:head>
<main>
  <h1>
    <img
      src={isISActive ? isLogo : etwLogo}
      title={pageTitle}
      alt={pageTitle}
    />
  </h1>
  Switch status: {isISActive ? "woop" : "doop"}
  <div class="buttonContainer">
    <Button text={"Play"} handleClick={handlePlay} />
    <Button text={"Switch"} handleClick={handleSwitch} />
    <Button text={"Website"} handleClick={handleWebsite} />
    <Button text={"Uninstall"} handleClick={handleUninstall} />
    <Button text={"Exit"} handleClick={handleExit} />
  </div>
  <footer>
    <span class="prefix">v</span><span class="version">{version}</span>
  </footer>
</main>

<style>
  main {
    --font-size-factor: calc(40 * calc((100vh - 800px) / (1200 - 800)));
    --font-size: clamp(32px, var(--font-size-factor), 40px);

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
    height: calc(var(--width-ratio) * 374);
    object-fit: contain;
  }

  .buttonContainer {
    flex: 1;
    display: flex;
    justify-content: space-evenly;
    align-items: flex-start;
  }

  footer {
    text-align: right;
    margin-bottom: -0.25rem;
    padding-right: 0.25rem;
    font-family: "IM FELL English";
  }

  footer > .prefix {
    font-size: 2rem;
  }

  footer > .version {
    font-size: var(--font-size);
  }
</style>
