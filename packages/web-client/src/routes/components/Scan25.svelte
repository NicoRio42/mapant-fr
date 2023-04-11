<script lang="ts">
	import type { Map } from 'ol';
	import TileLayer from 'ol/layer/Tile';
	import { getContext, onDestroy, onMount } from 'svelte';
	import WMTS from 'ol/source/WMTS.js';
	import WMTSTileGrid from 'ol/tilegrid/WMTS.js';

	const getMap = getContext<() => Map>('map');
	let map: Map;
	let tileLayer: TileLayer<WMTS>;

	onMount(() => {
		map = getMap();
		const resolutions = [
			156543.03392804103, 78271.5169640205, 39135.75848201024, 19567.879241005125,
			9783.939620502562, 4891.969810251281, 2445.9849051256406, 1222.9924525628203,
			611.4962262814101, 305.74811314070485, 152.87405657035254, 76.43702828517625,
			38.218514142588134, 19.109257071294063, 9.554628535647034, 4.777314267823517,
			2.3886571339117584, 1.1943285669558792, 0.5971642834779396, 0.29858214173896974,
			0.14929107086948493, 0.07464553543474241
		];

		tileLayer = new TileLayer({
			source: new WMTS({
				url: 'https://wxs.ign.fr/zgof7z935ruzrv2m1obyjvq8/geoportail/wmts',
				layer: 'GEOGRAPHICALGRIDSYSTEMS.MAPS.SCAN25TOUR.L93',
				matrixSet: 'LAMB93',
				format: 'image/jpeg',
				style: 'normal',
				tileGrid: new WMTSTileGrid({
					origin: [0, 0], // topLeftCorner
					resolutions, // résolutions
					matrixIds: [
						'0',
						'1',
						'2',
						'3',
						'4',
						'5',
						'6',
						'7',
						'8',
						'9',
						'10',
						'11',
						'12',
						'13',
						'14',
						'15',
						'16',
						'17',
						'18',
						'19'
					] // ids des TileMatrix
				})
			}),
			opacity: 0.7
		});

		map?.addLayer(tileLayer);
	});

	onDestroy(() => {
		if (map !== undefined && tileLayer !== undefined) map.removeLayer(tileLayer);
	});
</script>
