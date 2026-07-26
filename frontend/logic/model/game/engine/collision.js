import {gameMap, layersForRoom} from "../storage/gameState.js";
import {PLAYER_LOCAL_POINTS} from "../storage/gameConstants.js";

const TILE_NORMALS = [{x: 0, y: -1},
    {x: 1, y: 0},
    {x: 0, y: 1},
    {x: -1, y: 0}];

export function canMoveTo(nextPosition, player) {
      const personPoints = getPlayerPoints(nextPosition, player.direction);
      const personNormals = getPlayerNormals(personPoints);
      const personSAT = {points: personPoints, normals: personNormals};
      const nearby = getNearbyTiles(nextPosition.x, nextPosition.y, gameMap[player.roomId]);
      for (const tile of nearby) {
          const tileInfo = layersForRoom.tilesInfo[tile.tileId];
          if (!tileInfo.blocksPlayer) continue;

          for (const hitbox of tileInfo.hitboxes) {
              const tilePoints = getTileWorldPoints(tile.x, tile.y, hitbox);
              const tileSAT = {points: tilePoints, normals: TILE_NORMALS};
              if (checkCollisionSAT(personSAT, tileSAT)) {
                  console.log('COLLISIONS:', tileSAT);
                  return false;
              }
          }
      }

    return true;
}

function getTileWorldPoints(col, row, hitbox) {
    const s = layersForRoom.tileSize;
    const x = col * s + hitbox.x;
    const y = row * s + hitbox.y;
    return [
        {x, y},
        {x: x + hitbox.w, y},
        {x: x + hitbox.w, y: y + hitbox.h},
        {x, y: y + hitbox.h}
    ];
}

function getPlayerPoints(centerPlayer, angle) {
    const cosP = Math.cos(angle);
    const sinP = Math.sin(angle);

    return PLAYER_LOCAL_POINTS.map(point => ({
        x: centerPlayer.x + (point.x * cosP - point.y * sinP),
        y: centerPlayer.y + (point.x * sinP + point.y * cosP)
    }));
}

function getPlayerNormals(points) {
    const normals = [];
    for (let i = 0; i < points.length; i++) {
        const p1 = points[i];
        const p2 = points[(i + 1) % points.length];
        const edge = {x: p2.x - p1.x, y: p2.y - p1.y};
        normals.push({x: -edge.y, y: edge.x});
    }
    return normals;
}

function checkCollisionSAT(box1, box2) {
    if (!checkAxes(box1.points, box2.points, box1.normals)) {
        return false;
    }
    return checkAxes(box1.points, box2.points, box2.normals);
}

function checkAxes(points1, points2, normals) {
    for (const axis of normals) {
        const proj1 = getMinMax(points1, axis);
        const proj2 = getMinMax(points2, axis);
        if (!isOverlapping(proj1, proj2)) {
            return false;
        }
    }
    return true;
}

function getMinMax(points, axis) {
    let min = Infinity;
    let max = -Infinity;
    for (const point of points) {
        const dot = point.x * axis.x + point.y * axis.y;
        if (dot < min) {
            min = dot;
        }
        if (dot > max) {
            max = dot;
        }
    }
    return {min, max};
}

function isOverlapping(projection1, projection2) {
    return (projection1.max >= projection2.min) && (projection2.max >= projection1.min);
}

function getNearbyTiles(x, y, mapData) {
    const center = getTileAtPosition(x, y);
    const map = mapData.collision;
    const mapWidth = layersForRoom.width;
    const mapHeight = layersForRoom.height;
    const nearby = [];
    for (let row = center.row - 1; row <= center.row + 1; row++) {
        for (let col = center.col - 1; col <= center.col + 1; col++) {
            if ((row >= 0) && (row < mapHeight) && (col >= 0) && (col < mapWidth)) {
                const index = row * mapWidth + col;
                const tileId = map.data[index];
                if (tileId > 0) {
                    nearby.push({
                        tileId: tileId - 1,
                        x: col,
                        y: row
                    });
                }
            }
        }
    }
    return nearby;
}

function getTileAtPosition(x, y) {
    const TILE_SIZE = layersForRoom.tileSize;

    const col = Math.floor(x / TILE_SIZE);
    const row = Math.floor(y / TILE_SIZE);
    return {row: row, col: col};
}