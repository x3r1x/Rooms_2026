import {GAME_CONSTANTS, GAME_SPRITES} from "../../model/game/gameConstants.js";

export function drawBullets(canvas, context, state) {
    state.bullets.forEach(function (bullet) {
        context.save();
        context.translate(bullet.x , bullet.y );
        context.rotate(bullet.direction + Math.PI / 2);
        const sprite = GAME_SPRITES.BULLET_FLIES;
        context.drawImage(sprite, -GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
            GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
        context.restore();
    })
}