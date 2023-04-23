<script lang="ts">
	import type { Map } from 'ol';
	import type { Type } from 'ol/geom/Geometry';
	import Draw, { DrawEvent, type GeometryFunction } from 'ol/interaction/Draw.js';
	import { createEventDispatcher, getContext, onDestroy, onMount } from 'svelte';

	export let type: Type;
	export let geometryFunction: GeometryFunction | undefined;

	const dispatch = createEventDispatcher<{ drawEnd: DrawEvent }>();
	let draw: Draw, map: Map;
	const getMap = getContext<() => Map>('map');

	function escapeCallback(event: KeyboardEvent) {
		if (event.code === 'Escape' && draw !== undefined) draw.abortDrawing();
	}

	onMount(() => {
		map = getMap();

		draw = new Draw({
			type,
			geometryFunction
		});

		draw.on('drawend', (e) => dispatch('drawEnd', e));
		map.addInteraction(draw);

		document.addEventListener('keydown', escapeCallback);
	});

	onDestroy(() => {
		if (map !== undefined && draw !== undefined) map.removeInteraction(draw);
		document.removeEventListener('keydown', escapeCallback);
	});
</script>
