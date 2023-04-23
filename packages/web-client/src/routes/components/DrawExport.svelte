<script>
	import { createBox } from 'ol/interaction/Draw';
	import Draw from './Draw.svelte';
	import { clientExport } from '$lib/client-export';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
</script>

<Draw
	type="Circle"
	geometryFunction={createBox()}
	on:drawEnd={(e) => {
		clientExport(e);
		dispatch('drawEnd');
	}}
/>
