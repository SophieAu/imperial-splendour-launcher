<script lang="ts">
  import * as styles from './styles.button';

  export let text: string;
  export let onClick: () => void;
</script>

<button class={styles.root} on:click={onClick}>{text}</button>
