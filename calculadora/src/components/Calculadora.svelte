<script>
    import Button from "./Button.svelte";

    const buttons = ['C', '/', '*', '-', '7', '8', '9', '+', '4', '5', '6', '1', '2', '3', '=', '0', '.']
</script>

<section class="border-2 rounded-lg p-2 pt-4 pb-4">
    <div class="h-14 bg-gray-400 m-2 mt-0 p-2 flex justify-end">0</div>
    <div class="grid grid-rows-5 grid-cols-4 gap-1">
        {#each buttons as button}
            <Button value={button} />
        {/each}
    </div>
</section>