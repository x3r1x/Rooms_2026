import {TILE_IMG} from "../../../model/game/storage/gameConstants.js";
import {layersForRoom, room, gameMap} from "../../../model/game/storage/gameState.js";

export function drawBackground(canvas, context) {
    console.log(gameMap[0]);
    //const mapExit = gameMap.exits.data;
    const mapFloor = gameMap[0].floors.data;
    const mapWall = gameMap[0].walls.data;
    const mapObject = gameMap[0].object.data;
    const mapWidth = layersForRoom.width;

    context.fillStyle = "#1f2535";
    context.fillRect(0, 0, canvas.width, canvas.height);

    drawLayer(context, mapFloor, mapWidth);
    //drawLayer(context, mapExit, mapWidth);
    drawLayer(context, mapWall, mapWidth);
    drawLayer(context, mapObject, mapWidth);
}
function drawLayer(context, layer, mapWidth) {
    const tileSize = 36;
    const tilesPerRow = 37;
    layer.forEach((tileId, index) => {
        if (tileId === 0) return;
        const canvasX = (index % mapWidth) * tileSize;
        const canvasY = Math.floor(index / mapWidth) * tileSize;

        const id = tileId - 1;
        const sourceX = (id % tilesPerRow) * tileSize;
        const sourceY = Math.floor(id / tilesPerRow) * tileSize;
        context.drawImage(
            TILE_IMG.TILE,
            sourceX, sourceY, tileSize, tileSize,
            + canvasX, canvasY, tileSize, tileSize
        );
    });
}