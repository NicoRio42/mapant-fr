import { transformExtent } from 'ol/proj';
import { TILES_BASE_URL } from '../environments/environment';
import { lat2tile, lon2tile, tile2lat, tile2long } from './helpers';
import type { DrawEvent } from 'ol/interaction/Draw';

const EXPORT_TILE_LIMIT = 20;
const TILE_PIXEL_SIZE = 591;
const ZOOM = 13;

export async function clientExport(e: CustomEvent<DrawEvent>) {
	const extent = e.detail.feature.getGeometry()?.getExtent();
	if (!extent) return;
	const [lon1, lat1, lon2, lat2] = transformExtent(extent, 'EPSG:3857', 'EPSG:4326');
	const xtile1 = lon2tile(lon1, ZOOM);
	const ytile1 = lat2tile(lat1, ZOOM);
	const xtile2 = lon2tile(lon2, ZOOM);
	const ytile2 = lat2tile(lat2, ZOOM);

	if (
		Math.abs(xtile1 - xtile2) > EXPORT_TILE_LIMIT ||
		Math.abs(ytile1 - ytile2) > EXPORT_TILE_LIMIT
	) {
		alert('Area is too big for an export');
		return;
	}

	const tiles: string[][] = [];
	const canvas = document.createElement('canvas');
	const ctx = canvas.getContext('2d')!;

	const minXTile = Math.min(xtile1, xtile2);
	const maxXTile = Math.max(xtile1, xtile2);
	const minYTile = Math.min(ytile1, ytile2);
	const maxYTile = Math.max(ytile1, ytile2);

	for (let x = minXTile; x <= maxXTile; x++) {
		const column: string[] = [];

		for (let y = minYTile; y <= maxYTile; y++) {
			column.push(`${TILES_BASE_URL}/${ZOOM}/${x}/${y}.png`);
		}

		tiles.push(column);
	}

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

	const maxX = Math.abs(xtile1 - xtile2);
	const maxY = Math.abs(ytile1 - ytile2);

	for (let x = 0; x <= maxX; x++) {
		for (let y = 0; y <= maxY; y++) {
			const blob = await blobs[x][y];
			const img = new Image();

			img.onload = () => {
				imgs.push({ img, x: x * TILE_PIXEL_SIZE, y: y * TILE_PIXEL_SIZE });
				imgsCount++;
				if (imgsCount === imgsNumber) drawImgsAndDownloadFile();
			};

			img.src = URL.createObjectURL(blob);
		}
	}
}
