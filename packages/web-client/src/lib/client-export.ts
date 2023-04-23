import { transformExtent } from 'ol/proj';
import { TILES_BASE_URL } from '../environments/environment';
import { lat2tile, lon2tile, tile2lat, tile2long } from './helpers';
import type { DrawEvent } from 'ol/interaction/Draw';

const EXPORT_TILE_LIMIT = 20;
const TILE_PIXEL_SIZE = 591;

export async function clientExport(e: CustomEvent<DrawEvent>) {
	const extent = e.detail.feature.getGeometry()?.getExtent();
	if (!extent) return;
	const zoom = 13;
	const [lon1, lat1, lon2, lat2] = transformExtent(extent, 'EPSG:3857', 'EPSG:4326');
	const xtile1 = lon2tile(lon1, zoom);
	const ytile1 = lat2tile(lat1, zoom);
	const xtile2 = lon2tile(lon2, zoom);
	const ytile2 = lat2tile(lat2, zoom);

	if (
		Math.abs(xtile1 - xtile2) > EXPORT_TILE_LIMIT ||
		Math.abs(ytile1 - ytile2) > EXPORT_TILE_LIMIT
	) {
		alert('Area is too big for an export');
		return;
	}

	const lon1Before = tile2long(xtile1, zoom);
	const lon1After = tile2long(xtile1 + 1, zoom);
	const offsetX1 = Math.abs((TILE_PIXEL_SIZE * (lon1 - lon1Before)) / (lon1After - lon1Before));

	const lat1Before = tile2lat(ytile1, zoom);
	const lat1After = tile2lat(ytile1 + 1, zoom);
	const offsetY1 = Math.abs((TILE_PIXEL_SIZE * (lat1 - lat1Before)) / (lat1After - lat1Before));

	const lon2Before = tile2long(xtile2, zoom);
	const lon2After = tile2long(xtile2 + 1, zoom);
	const offsetX2 = Math.abs((TILE_PIXEL_SIZE * (lon2 - lon2Before)) / (lon2After - lon2Before));

	const lat2Before = tile2lat(ytile2, zoom);
	const lat2After = tile2lat(ytile2 + 1, zoom);
	const offsetY2 = Math.abs((TILE_PIXEL_SIZE * (lat2 - lat2Before)) / (lat2After - lat2Before));

	const tiles: string[][] = [];
	const canvas = document.createElement('canvas');
	const ctx = canvas.getContext('2d')!;

	for (let x = Math.min(xtile1, xtile2); x <= Math.max(xtile1, xtile2); x++) {
		const column: string[] = [];

		for (let y = Math.min(ytile1, ytile2); y <= Math.max(ytile1, ytile2); y++) {
			column.push(`${TILES_BASE_URL}/${zoom}/${x}/${y}.png`);
		}

		tiles.push(column);
	}

	// canvas.width = TILE_PIXEL_SIZE * tiles.length - offsetX1 + offsetX2;
	// canvas.height = TILE_PIXEL_SIZE * tiles[0].length - offsetY1 + offsetY2;

	canvas.width = TILE_PIXEL_SIZE * tiles.length;
	canvas.height = TILE_PIXEL_SIZE * tiles[0].length;

	const blobs = await Promise.all(
		tiles.map((column) => column.map((url) => fetch(url).then((r) => r.blob())))
	);

	const imgs: { img: HTMLImageElement; x: number; y: number }[] = [];
	let imgsCount = 0;
	const imgsNumber = tiles.length * tiles[0].length;

	function drawImgsAndDownloadFile() {
		imgs.forEach((img) => ctx.drawImage(img.img, img.x, img.y));

		const link = document.createElement('a');
		link.download = 'mapant-fr-export.png';
		link.href = canvas.toDataURL();
		link.click();
	}

	for (let x = 0; x <= Math.abs(xtile1 - xtile2); x++) {
		for (let y = 0; y <= Math.abs(ytile1 - ytile2); y++) {
			const blob = await blobs[x][y];
			const img = new Image();
			img.onload = () => {
				imgs.push({ img, x: x * TILE_PIXEL_SIZE, y: y * TILE_PIXEL_SIZE });
				// imgs.push({ img, x: x * TILE_PIXEL_SIZE - offsetX1, y: y * TILE_PIXEL_SIZE - offsetY1 });
				imgsCount++;
				if (imgsCount === imgsNumber) drawImgsAndDownloadFile();
			};

			img.src = URL.createObjectURL(blob);
		}
	}
}
