<script>
  import Button from './Button.svelte';

  // export let message;
  export let onClick;

  const buttonText = 'OK';
</script>

<div class="overlay">
  <div class="modal">
    <p class="message">
      We couldn't cleanly uninstall Imperial Splendour. Please manully delete any remaining files.
    </p>
    <div class="modalButton">
      <Button text={buttonText} handleClick={onClick} />
    </div>
  </div>
</div>

<style>
  .overlay {
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    display: grid;
    place-items: center center;

    background: rgba(255, 255, 255, 0.3);
    backdrop-filter: blur(2px);
  }

  .modal {
    width: 33vw;
    background: var(--modal-bg);

    border: 0.25rem solid rgb(77, 50, 50);
    border-radius: 0.375rem;
    padding: 0.75rem;

    display: flex;
    flex-direction: column;
  }
  .modalButton {
    align-self: flex-end;
  }

  .message {
    font: normal var(--font-size-body) var(--font-family-body);

    margin: 0;
    padding: 0.5rem 0.5rem 0.75rem;
  }
</style>
