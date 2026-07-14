import {TILE_IMG} from "../../model/storage/gameConstants.js";
import {lastState, layersForRoom, room} from "../../model/game/gameModel.js";

const x = 1;

export function drawBackground(canvas, context, state) {
    const map = room.exits.data;
    const mapFloor = room.floors.data;
    const mapWall = room.walls.data;
    const mapObject = room.object.data;
    const tileSize = 36;
    const tilesPerRow = 37;
    const mapWidth = layersForRoom.width;
    context.fillStyle = "#1f2535";
    context.fillRect(0, 0, canvas.width, canvas.height);
    mapFloor.forEach((tileId, index) => {
        if (tileId === 0) return;
        const canvasX = (index % mapWidth) * x * tileSize;
        const canvasY = Math.floor(index / mapWidth) * x * tileSize;

        const id = tileId - 1;
        const sourceX = (id % tilesPerRow) * tileSize;
        const sourceY = Math.floor(id / tilesPerRow) * tileSize;
        context.drawImage(
            TILE_IMG.TILE,
            sourceX, sourceY, tileSize, tileSize,
            200 + canvasX, canvasY, x * tileSize, x * tileSize
        );
    });
    map.forEach((tileId, index) => {
        if (tileId === 0) return;
        const canvasX = (index % mapWidth) * x * tileSize;
        const canvasY = Math.floor(index / mapWidth) * x * tileSize;

        const id = tileId - 1;
        const sourceX = (id % tilesPerRow) * tileSize;
        const sourceY = Math.floor(id / tilesPerRow) * tileSize;
        context.drawImage(
            TILE_IMG.TILE,
            sourceX, sourceY, tileSize, tileSize,
            200 + canvasX, canvasY, x * tileSize, x * tileSize
        );
    });

    mapWall.forEach((tileId, index) => {
        if (tileId === 0) return;
        const canvasX = (index % mapWidth) * x * tileSize;
        const canvasY = Math.floor(index / mapWidth) * x * tileSize;

        const id = tileId - 1;
        const sourceX = (id % tilesPerRow) * tileSize;
        const sourceY = Math.floor(id / tilesPerRow) * tileSize;
        context.drawImage(
            TILE_IMG.TILE,
            sourceX, sourceY, tileSize, tileSize,
            200 + canvasX, canvasY, x * tileSize, x * tileSize
        );
    });
    mapObject.forEach((tileId, index) => {
        if (tileId === 0) return;
        const canvasX = (index % mapWidth) * x * tileSize;
        const canvasY = Math.floor(index / mapWidth) * x * tileSize;

        const id = tileId - 1;
        const sourceX = (id % tilesPerRow) * tileSize;
        const sourceY = Math.floor(id / tilesPerRow) * tileSize;
        context.drawImage(
            TILE_IMG.TILE,
            sourceX, sourceY, tileSize, tileSize,
            200 + canvasX, canvasY, x * tileSize, x * tileSize
        );
    });
    //TODO вынести отрисовку матрицы в отдельную функцию
    //TODO подтягивать ширину матрицы из JSON файла
}
function drawLayer(context, state, layer) {

}