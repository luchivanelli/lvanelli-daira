<script>
    export let value
</script>

{#if value == '+' || value == '='}
    <button on:click={$$props.handleDiv} class="row-span-2 text-lg">{value}</button>
{:else if value == '0'}
    <button on:click={$$props.handleDiv} class="col-span-2 text-lg">{value}</button>
{:else}
    <button on:click={$$props.handleDiv}>{value}</button>
{/if}