import {GAME_CONSTANTS} from "../gameConstants.js";

export function drawBullets(canvas, context, state) {
    state.bullets.forEach(function (bullet) {
        context.save();
        context.translate(bullet.x + GAME_CONSTANTS.BULLET_WIDTH / 2, bullet.y + GAME_CONSTANTS.BULLET_HEIGHT / 2);
        context.rotate(bullet.direction + Math.PI / 2);
        context.fillStyle = GAME_CONSTANTS.BULLET_COLOR;
        context.fillRect(-GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
            GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
        context.restore();
    })
}