<script lang="ts">
    // @ts-nocheck
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

    const handleDiv = e => {
        let value = e.detail;
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
    };

    const operation = () => {
        let newValue, resultado;

        if (valueDiv.includes('+')) {
            //verificar el operador
            newValue = valueDiv.split('+'); //separar string, obtener numeros a operar
            newValue[1] = newValue[1].replace('=', ''); //eliminar '=' del 2do numero
            resultado = parseInt(newValue[0]) + parseInt(newValue[1]); //realizar operacion
        } else if (valueDiv.includes('-')) {
            newValue = valueDiv.split('-');
            newValue[1] = newValue[1].replace('=', '');
            resultado = parseInt(newValue[0]) - parseInt(newValue[1]);
        } else if (valueDiv.includes('*')) {
            newValue = valueDiv.split('*');
            newValue[1] = newValue[1].replace('=', '');
            resultado = parseInt(newValue[0]) * parseInt(newValue[1]);
        } else if (valueDiv.includes('/')) {
            newValue = valueDiv.split('/');
            newValue[1] = newValue[1].replace('=', '');
            resultado = parseInt(newValue[0]) / parseInt(newValue[1]);
        } else if (valueDiv.includes('%')) {
            newValue = valueDiv.split('%');
            newValue[1] = newValue[1].replace('=', '');
            resultado = parseInt(newValue[0]) / parseInt(newValue[1]);
        }

        resultado = valueDiv = resultado.toString();
    };
</script>

<section class="container p-4 pt-0 bg-red-100 min-h-screen grid place-content-center">
    <p class="text-center p-2 mb-4">Calculadora Basica</p>
    <div class="display p-2 bg-black text-white text-end">{valueDiv}</div>
    <div class="buttons grid grid-cols-4 grid-rows-5">
        {#each buttons as button}
            <Button value={button} on:enviar={handleDiv} />
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
