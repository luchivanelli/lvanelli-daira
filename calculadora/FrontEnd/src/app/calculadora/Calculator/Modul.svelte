<script lang="ts">
    import axios from 'axios';
    import Button from './components/Button.svelte';

    const buttons = ['CE', 'C', '±', '%', 7, 8, 9, '/', 4, 5, 6, '*', 1, 2, 3, '-', 0, '.', '=', '+'];
    let valueDiv = '0';

    const getData = async () => {
        try {
            const response = await axios.get('http://localhost:5173/operations');
            console.log(response);
        } catch (err) {
            console.log('error');
        }
    };

    getData();

    // @ts-ignore
    const handleDiv = e => {
        let value = e.target.innerHTML;
        if (valueDiv == '0') {
            //limpiar el '0' inicial
            //validar que el primer caracter sea un numero o un '-' (para operar numero negativo)
            if (!isNaN(parseInt(value)) || value == '-') {
                valueDiv = '';
                valueDiv = valueDiv + value;
            }
        } else {
            valueDiv = valueDiv + value; //concatenar
        }

        if (value == 'C') {
            //limpiar display
            valueDiv = '0';
        } else if (value == '=') {
            //operar
            operation();
        }

        //hacer varias operaciones a la vez
        if (isNaN(parseInt(value))) {
            //evaluar si value no es un numero
            let newDiv = valueDiv.slice(1, valueDiv.length - 1); //newDiv: contiene el valor de valueDiv sin el primer y ultimo caracter
            newDiv.includes('+') || newDiv.includes('-') || newDiv.includes('/') || newDiv.includes('*') ? operation(value, valueDiv.slice(0, valueDiv.length - 1)) : null;
            //si ya hay una operacion para hacer, se llama a la funcion.
            //luego se agrega el operador correspondiente (value) al resultado de la operacion
        }
    };

    // @ts-ignore
    const operation = (value, newDiv) => {
        let newValue, resultado;
        newDiv ? (valueDiv = newDiv) : null; //newDiv contiene el valor del display menos el ultimo caracter

        if (valueDiv.includes('+')) {
            //verificar el operador
            newValue = valueDiv.split('+'); //separar string, obtener numeros a operar
            // @ts-ignore
            newValue[1] = newValue[1].replace('=', ''); //eliminar '=' del 2do numero
            // @ts-ignore
            resultado = parseFloat(newValue[0]) + parseFloat(newValue[1]); //realizar operacion
        } else if (valueDiv.includes('-')) {
            newValue = valueDiv.split('-');
            // @ts-ignore
            newValue[1] = newValue[1].replace('=', '');
            // @ts-ignore
            resultado = parseFloat(newValue[0]) - parseFloat(newValue[1]);
        } else if (valueDiv.includes('*')) {
            newValue = valueDiv.split('*');
            // @ts-ignore
            newValue[1] = newValue[1].replace('=', '');
            // @ts-ignore
            resultado = parseFloat(newValue[0]) * parseFloat(newValue[1]);
        } else if (valueDiv.includes('/')) {
            newValue = valueDiv.split('/');
            // @ts-ignore
            newValue[1] = newValue[1].replace('=', '');
            // @ts-ignore
            resultado = parseFloat(newValue[0]) / parseFloat(newValue[1]);
        }

        // @ts-ignore
        resultado % 1 == 0 ? (valueDiv = resultado.toString()) : (valueDiv = resultado.toFixed(3).toString());
        //si el numero es flotante, limitar la cantidad de decimales. Pasar variable de number a string
        value ? (valueDiv = valueDiv + value) : null; //agregamos operador
    };
</script>

<section class="container p-4 pt-0">
    <p class="text-center p-2 mb-4">Calculadora Basica</p>
    <div class="display p-2 bg-black text-white text-end">{valueDiv}</div>
    <div class="buttons grid grid-cols-4 grid-rows-5">
        {#each buttons as button}
            <Button value={button} {handleDiv} />
        {/each}
    </div>
</section>

<style>
    p {
        font-size: 22px;
    }
    .container {
        border: 4px solid black;
        /* padding: 4px; */
        width: 250px;
        margin: 100px auto;
    }

    .display {
        height: 50px;
        text-align: end;
        font-size: 26px;
        margin-bottom: 4px;
    }

    .buttons {
        gap: 3px;
        grid-template-columns: repeat(4, minmax(0, 50px));
        grid-template-rows: repeat(5, minmax(0, 50px));
    }
</style>
