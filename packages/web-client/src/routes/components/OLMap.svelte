<script lang="ts">
	import { Map, View } from 'ol';
	import 'ol/ol.css';
	import { onDestroy, onMount, setContext } from 'svelte';
	import DoubleClickZoom from 'ol/interaction/DoubleClickZoom.js';
	import proj4 from 'proj4';
	import { get as getProjection } from 'ol/proj.js';
	import { register } from 'ol/proj/proj4.js';
	import { transform } from 'ol/proj';

	// export let angle: number;
	// export let fitBox: [number, number, number, number];
	// export let padding: [number, number, number, number];

	let map: Map;
	let view: View;

	// $: {
	// 	if (view !== undefined) {
	// 		view.setRotation(angle);
	// 		view.fit(fitBox, { padding });
	// 	}
	// }

	setContext('map', () => map);

	onMount(() => {
		// proj4.defs(
		// 	'EPSG:2154',
		// 	'+proj=lcc +lat_0=46.5 +lon_0=3 +lat_1=49 +lat_2=44 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m +no_defs +type=crs'
		// );

		// register(proj4);
		// const proj2154 = getProjection('EPSG:2154');
		// proj2154?.setExtent([-378305.81, 6005281.2, 1320649.57, 7235612.72]);

		const center = transform([45.1344, -14.2355], 'EPSG:4326', 'EPSG:3857');

		map = new Map({
			target: 'mapviewer',
			view: new View({
				center,
				zoom: 6,
				minZoom: 0,
				maxZoom: 15
				// projection: 'EPSG:2154'
			})
		});

		map.addInteraction(new DoubleClickZoom());

		// view.fit(fitBox, { padding });
	});

	onDestroy(() => {
		if (map !== undefined) map.dispose();
	});
</script>

<div id="mapviewer" class="map" />

{#if map}
	<slot />
{/if}

<style>
	#mapviewer {
		width: 100%;
		height: 100%;
		cursor: grab;
	}

	#mapviewer:active {
		cursor: grabbing;
	}
</style>
