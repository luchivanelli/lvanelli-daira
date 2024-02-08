<script>
    // @ts-nocheck

    export let value;
</script>

{#if value == 'CE' || value == 'C'}
    <button on:click={$$props.handleDiv} class="clean bg-orange-600 text-center">{value}</button>
{:else if value == '='}
    <button on:click={$$props.handleDiv} class="equal bg-green-600 text-center">{value}</button>
{:else if !isNaN(value) || value == '.'}
    <button on:click={$$props.handleDiv} class="number bg-blue-600 text-center">{value}</button>
{:else}
    <button on:click={$$props.handleDiv} class="operator bg-red-900 text-center">{value}</button>
{/if}

<style>
    button {
        color: white;
    }
    .clean {
        background-color: orangered;
    }
    .equal {
        background-color: green;
    }
    .number {
        background-color: blue;
    }
    .operator {
        background-color: rgb(116, 18, 18);
    }
</style>
