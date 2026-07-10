import {GAME_CONSTANTS, TILE_IMG} from "../../model/game/gameConstants.js";

export function drawBackground(canvas, context, state) {
    const map = state.map;
    const map2 = state.mapCollisian;
    const tileSize = 32;
    const tilesPerRow = 37;
    const mapWidth = 35;
    context.fillStyle = "#000";
    context.fillRect(0, 0, canvas.width, canvas.height);
    map.forEach((tileId, index) => {
        if (tileId === 0) return;
        const canvasX = (index % mapWidth) * tileSize;
        const canvasY = Math.floor(index / mapWidth) * tileSize;

        const id = tileId - 1;
        const sourceX = (id % tilesPerRow) * tileSize;
        const sourceY = Math.floor(id / tilesPerRow) * tileSize;
        context.drawImage(
            TILE_IMG.TILE,
            sourceX, sourceY, tileSize, tileSize,
            200+canvasX, canvasY, tileSize, tileSize
        );
    });
    map2.forEach((tileId, index) => {
        if (tileId === 0) return;
        const canvasX = (index % mapWidth) * tileSize;
        const canvasY = Math.floor(index / mapWidth) * tileSize;

        const id = tileId - 1;
        const sourceX = (id % tilesPerRow) * tileSize;
        const sourceY = Math.floor(id / tilesPerRow) * tileSize;
        context.drawImage(
            TILE_IMG.TILE,
            sourceX, sourceY, tileSize, tileSize,
            200+canvasX, canvasY, tileSize, tileSize
        );
    });
 //TODO вынести отрисовку матрицы в отдельную функцию
 //TODO подтягивать ширину матрицы из JSON файла
}