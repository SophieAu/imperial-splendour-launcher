<script lang="ts">
  import Button from './Button.svelte';
  import { modalButtonText } from './strings';
  import * as styles from './styles.modal';

  export let onClick: () => void;
</script>

<div class={styles.overlay}>
  <div class={styles.container}>
    <p class={styles.message}>
      <slot />
    </p>
    <div class={styles.button}>
      <Button text={modalButtonText} {onClick} />
    </div>
  </div>
</div>
