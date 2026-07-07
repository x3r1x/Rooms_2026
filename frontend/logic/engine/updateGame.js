import {keys} from './listeners.js';
import {updateBullets} from "./bullet.js";
import {updatePlayer} from "./player.js";

export function updateGame(elapsedTime, state) {
    const playerDirection = {
        x: 0,
        y: 0
    }

    updateBullets(elapsedTime, state)
    updateDirection(playerDirection);
    updatePlayer(playerDirection, elapsedTime, state);
}

function updateDirection(direction) {
    if (keys['w'] || keys['ц'] || keys['arrowup']) {
        direction.y -= 1
    }
    if (keys['s'] || keys['ы'] || keys['arrowdown']) {
        direction.y += 1
    }
    if (keys['a'] || keys['ф'] || keys['arrowleft']) {
        direction.x -= 1
    }
    if (keys['d'] || keys['в'] || keys['arrowright']) {
        direction.x += 1
    }

    let vectorLength = (Math.sqrt(direction.x * direction.x + direction.y * direction.y));
    if (vectorLength > 0) {
        direction.x /= vectorLength;
        direction.y /= vectorLength;
    }
}