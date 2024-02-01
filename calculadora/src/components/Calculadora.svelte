<script>
    import Button from "./Button.svelte";

    const buttons = ['C', '/', '*', '-', 7, 8, 9, '+', 4, 5, 6, 1, 2, 3, '=', 0, '.']
    let valueDiv = '0'
    let a, b

    const handleDiv = (e)=> {
        let value = e.target.innerHTML
        if (valueDiv == '0') {
            valueDiv = ''
            valueDiv = valueDiv + value
        } else {
            valueDiv = valueDiv + value
        }

        if (value == 'C') {
            valueDiv = '0'
        }

        if (value == '=') {
            operation()
        }
        
    }

    const operation = ()=> {
    }

</script>

<section class="border-2 rounded-lg p-2 pt-4 pb-4">
    <div class="h-16 bg-gray-400 text-black font-bold m-2 mt-0 p-2 flex justify-end text-4xl">{valueDiv}</div>
    <div class="grid grid-rows-5 grid-cols-4 gap-1">
        {#each buttons as button}
            <Button value={button} handleDiv={handleDiv}/>
        {/each}
    </div>
</section>