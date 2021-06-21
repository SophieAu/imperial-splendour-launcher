<script lang="ts">
  import { modalButtonText } from '../strings';

  export let message: string;
  export let onClick: () => void;

</script>

<div class="overlay">
  <div class="container">
    <p class="message">
      {message}
    </p>
    <div class="button">
      <button on:click={onClick}>{modalButtonText}</button>
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
    display: flex;
    align-items: center;
    justify-content: center;

    background: rgba(255, 255, 255, 0.3);
    backdrop-filter: blur(2px);
  }

  .container {
    width: 50vw;

    border: 0.125rem solid rgb(77, 50, 50);
    border-radius: 0.375rem;
    padding: 0.75rem;

    display: flex;
    flex-direction: column;
  }

  .button {
    align-self: flex-end;
  }

  .message {
    margin: 0;
    padding: 0.5rem 0.5rem 0.75rem;
  }

</style>
