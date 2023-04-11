<script lang="ts">
	import type { Map } from 'ol';
	import TileLayer from 'ol/layer/Tile';
	import XYZ from 'ol/source/XYZ';
	import { getContext, onDestroy, onMount } from 'svelte';
	import { TILES_BASE_URL } from '../../environments/environment';

	const getMap = getContext<() => Map>('map');
	let map: Map;
	let tileLayer: TileLayer<XYZ>;

	onMount(() => {
		map = getMap();

		const url = `${TILES_BASE_URL}{z}/{x}/{y}.png`;

		tileLayer = new TileLayer({
			// extent: [343646, 1704354, 5619537, 7667537],
			source: new XYZ({
				url
			})
		});

		map?.addLayer(tileLayer);
	});

	onDestroy(() => {
		if (map !== undefined && tileLayer !== undefined) map.removeLayer(tileLayer);
	});
</script>
