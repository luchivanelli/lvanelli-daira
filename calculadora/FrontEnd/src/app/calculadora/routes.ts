
// @ts-nocheck
import { wrap } from 'svelte-spa-router/wrap'

const routes = { '/calculator': wrap({ asyncComponent: () => import('./Calculator/Modul.svelte') }), };

export default routes;
