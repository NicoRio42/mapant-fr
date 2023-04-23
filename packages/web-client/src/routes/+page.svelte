<script lang="ts">
	import Mapant from './components/Mapant.svelte';
	import OlMap from './components/OLMap.svelte';

	let isDrawingExport = false;
	let lazyDrawExport: Promise<typeof import('./components/DrawExport.svelte')>;

	function handleLoadSplitsClick() {
		isDrawingExport = !isDrawingExport;
		if (lazyDrawExport === undefined) lazyDrawExport = import('./components/DrawExport.svelte');
	}
</script>

<div class="wrapper" class:drawing={isDrawingExport}>
	<OlMap>
		<!-- <Osm /> -->
		<!-- <Scan25 /> -->

		<Mapant />

		{#if isDrawingExport && lazyDrawExport !== undefined}
			{#await lazyDrawExport then { default: DrawExport }}
				<DrawExport on:drawEnd={() => (isDrawingExport = false)} />
			{/await}
		{/if}
	</OlMap>

	<button type="button" class="export-btn" on:click={handleLoadSplitsClick}>
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="export-icon">
			<path
				d="M288 32c0-17.7-14.3-32-32-32s-32 14.3-32 32V274.7l-73.4-73.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l128 128c12.5 12.5 32.8 12.5 45.3 0l128-128c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L288 274.7V32zM64 352c-35.3 0-64 28.7-64 64v32c0 35.3 28.7 64 64 64H448c35.3 0 64-28.7 64-64V416c0-35.3-28.7-64-64-64H346.5l-45.3 45.3c-25 25-65.5 25-90.5 0L165.5 352H64zm368 56a24 24 0 1 1 0 48 24 24 0 1 1 0-48z"
			/>
		</svg>
	</button>
</div>

<style>
	.wrapper {
		position: relative;
		/* flex-shrink: 0; */
		flex-grow: 1;
		cursor: grab;
	}

	.wrapper:active {
		cursor: grabbing;
	}

	.wrapper.drawing {
		cursor: crosshair;
	}

	.export-btn {
		position: absolute;
		top: 1rem;
		right: 1rem;
		width: fit-content;
		backdrop-filter: blur(0.5rem);
		background-color: rgba(0, 0, 0, 0.486);
		border: none;
		padding: 0.25rem 0.5rem;
	}
	.export-icon {
		width: 1rem;
		height: 1rem;
	}

	@media screen and (max-width: 800px) {
		.export-btn {
			display: none;
		}
	}
</style>
