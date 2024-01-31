<script>
    export let value
</script>

{#if value == '+' || value == '='}
    <button class="row-span-2">{value}</button>
{:else if value == '0'}
    <button class="col-span-2">{value}</button>
{:else}
    <button>{value}</button>
{/if}