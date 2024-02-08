<script>
    // @ts-nocheck

    export let value;
    import { createEventDispatcher } from 'svelte';

    let dispatch = createEventDispatcher();

    const functionDispatch = () => {
        dispatch('enviar', value);
    };
</script>

{#if value == 'CE' || value == 'C'}
    <button on:click={functionDispatch} class="clean bg-orange-600 text-center">{value}</button>
{:else if value == '='}
    <button on:click={functionDispatch} class="equal bg-green-600 text-center">{value}</button>
{:else if !isNaN(value) || value == '.'}
    <button on:click={functionDispatch} class="number bg-blue-600 text-center">{value}</button>
{:else}
    <button on:click={functionDispatch} class="operator bg-red-900 text-center">{value}</button>
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
