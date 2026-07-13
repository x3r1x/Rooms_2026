import {GAME_CONSTANTS, GAME_SPRITES} from "../../model/game/gameConstants.js";

export function drawBullets(context, state) {
    for (const bullet of Object.values(state.player.bullets)) {
        drawPLayerBullets(context, bullet);
    }

    for (const enemy of Object.values(state.enemies)) {
        for (const enemyBullet of Object.values(enemy.bullets)) {
            drawEnemyBullets(context, enemyBullet);
        }
    }
}

function drawPLayerBullets(context, bullet) {
    const sprite = GAME_SPRITES.BULLET_FLIES;

    context.save();
    context.translate(bullet.x, bullet.y);
    context.rotate(bullet.direction + Math.PI / 2);
    context.drawImage(sprite, -GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
        GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
    context.restore();
}

function drawEnemyBullets(context, enemyBullet) {
    const sprite = GAME_SPRITES.ENEMY_BULLET_FLIES;

    context.save();
    context.translate(enemyBullet.x, enemyBullet.y);
    context.rotate(enemyBullet.direction + Math.PI / 2);
    context.drawImage(sprite, -GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
        GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
    context.restore();
}