import {canvas} from "../../main.js";
import {GAME_CONSTANTS} from "../../model/game/gameConstants.js";

export function updateBullets(elapsedTime, bulletsList) {
    const newBulletsList = {};

    for (const [id, bullet] of Object.entries(bulletsList)) {
        bullet.x += Math.cos(bullet.movementDirection) * elapsedTime * GAME_CONSTANTS.BULLET_SPEED;
        bullet.y += Math.sin(bullet.movementDirection) * elapsedTime * GAME_CONSTANTS.BULLET_SPEED;

        if (!isBulletOutOfBounds(bullet)) {
            newBulletsList[id] = bullet;
        }
    }

    return newBulletsList;
}

function isBulletOutOfBounds(bullet) {
    const didPassLeft = bullet.x + GAME_CONSTANTS.BULLET_WIDTH < 0;
    const didPassRight = bullet.x > canvas.width;
    const didPassTop = bullet.y + GAME_CONSTANTS.BULLET_HEIGHT < 0;
    const didPassBottom = bullet.y > canvas.height;

    return didPassLeft || didPassRight || didPassTop || didPassBottom;
}