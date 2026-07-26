import {GAME_CONSTANTS} from "../storage/gameConstants.js";
import {canMoveTo} from "./collision.js";
import {lerp, lerpDirection} from "../storage/interpolation.js";

export function updatePlayer(direction, elapsedTime, player) {
    updateVisualDirection(player);

    const nextX = player.x + direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    const nextY = player.y + direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;

    if (canMoveTo({x: nextX, y: nextY}, player)){
        player.x = nextX;
        player.y = nextY;
    }
}

export function updateVisualDirection(player) {
    player.direction = Math.atan2(player.mousePosition.y - player.y, player.mousePosition.x - player.x);
}

export function updateEnemies(dt, enemies, snaps) {
    snaps.stateA.p.forEach((playerStart) => {
        if (playerStart.id in enemies) {
            if (!(snaps.stateB.p.find((playerEnd) => playerEnd.id === playerStart.id))) {
                delete enemies[playerStart.id];
            } else {
                const playerEnd = snaps.stateB.p.find((playerEnd) => playerEnd.id === playerStart.id);

                enemies[playerStart.id].x = lerp(playerStart.x, playerEnd.x, dt);
                enemies[playerStart.id].y = lerp(playerStart.y, playerEnd.y, dt);
                enemies[playerStart.id].direction = lerpDirection(playerStart.a, playerEnd.a, dt);
                enemies[playerStart.id].hp = playerStart.h;
            }
        }
    })
}