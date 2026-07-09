import {canvas} from "../../main.js";
import {GAME_CONSTANTS} from "../../model/gameConstants.js";

export function updateBullets(elapsedTime, state) {
    state.bullets = state.bullets.filter(bullet => {
        bullet.x += Math.cos(bullet.direction) * elapsedTime * GAME_CONSTANTS.BULLET_SPEED;
        bullet.y += Math.sin(bullet.direction) * elapsedTime * GAME_CONSTANTS.BULLET_SPEED;

        return !isBulletOutOfBounds(bullet);
    })
}

function isBulletOutOfBounds(bullet) {
    const didPassLeft = bullet.x + GAME_CONSTANTS.BULLET_WIDTH < 0;
    const didPassRight = bullet.x > canvas.width;
    const didPassTop = bullet.y + GAME_CONSTANTS.BULLET_HEIGHT < 0;
    const didPassBottom = bullet.y > canvas.height;

    return didPassLeft || didPassRight || didPassTop || didPassBottom;
}