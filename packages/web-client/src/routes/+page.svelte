<script lang="ts">
	import Mapant from './components/Mapant.svelte';
	import OlMap from './components/OLMap.svelte';
	import Osm from './components/OSM.svelte';
	import Scan25 from './components/Scan25.svelte';
</script>

<div class="wrapper">
	<OlMap>
		<!-- <Osm /> -->
		<!-- <Scan25 /> -->

		<Mapant />
	</OlMap>
</div>

<style>
	.wrapper {
		position: relative;
		/* flex-shrink: 0; */
		flex-grow: 1;
	}
</style>
