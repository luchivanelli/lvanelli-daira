<script>
    import Button from "./Button.svelte";

    const buttons = ['C', '/', '*', '-', 7, 8, 9, '+', 4, 5, 6, 1, 2, 3, '=', 0, '.']
    let valueDiv = '0'

    const handleDiv = (e)=> {
        let value = e.target.innerHTML
        if (valueDiv == '0') {  //limpiar el '0' inicial
            //validar que el primer caracter sea un numero o un '-' (para operar numero negativo)
            if (!isNaN(parseInt(value)) || value == '-') {  
                valueDiv = ''
                valueDiv = valueDiv + value 
            }
        } else {
            valueDiv = valueDiv + value  //concatenar
        }

        if (value == 'C') {  //limpiar display
            valueDiv = '0'
        } else if (value == '=') {  //operar
            operation()
        } 
        
        //hacer varias operaciones a la vez
        if (isNaN(parseInt(value))) {  //evaluar si value no es un numero
            let newDiv = valueDiv.slice(1, valueDiv.length-1) //newDiv: contiene el valor de valueDiv sin el primer y ultimo caracter
            newDiv.includes('+') || 
            newDiv.includes('-') || 
            newDiv.includes('/') || 
            newDiv.includes('*') ? operation(value, valueDiv.slice(0, valueDiv.length-1)) : null 
            //si ya hay una operacion para hacer, se llama a la funcion.
            //luego se agrega el operador correspondiente (value) al resultado de la operacion
        }
    }

    const operation = (value, newDiv)=> {
        let newValue, resultado
        newDiv ? valueDiv = newDiv : null  //newDiv contiene el valor del display menos el ultimo caracter

        if (valueDiv.includes('+')) {  //verificar el operador
            newValue = valueDiv.split('+')  //separar string, obtener numeros a operar
            newValue[1] = newValue[1].replace('=', '')  //eliminar '=' del 2do numero
            resultado = parseFloat(newValue[0]) + parseFloat(newValue[1])  //realizar operacion

        } else if (valueDiv.includes('-')) {
            newValue = valueDiv.split('-')
            newValue[1] = newValue[1].replace('=', '')
            resultado = parseFloat(newValue[0]) - parseFloat(newValue[1])

        } else if (valueDiv.includes('*')) {
            newValue = valueDiv.split('*')
            newValue[1] = newValue[1].replace('=', '')
            resultado = parseFloat(newValue[0]) * parseFloat(newValue[1])

        } else if (valueDiv.includes('/')) {
            newValue = valueDiv.split('/')
            newValue[1] = newValue[1].replace('=', '')
            resultado = parseFloat(newValue[0]) / parseFloat(newValue[1])

        }
    
        resultado % 1 == 0 ? valueDiv = resultado.toString() : valueDiv = resultado.toFixed(3).toString()
        //si el numero es flotante, limitar la cantidad de decimales. Pasar variable de number a string
        value ? valueDiv = valueDiv + value : null  //agregamos operador
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