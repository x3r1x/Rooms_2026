import {GAME_CONSTANTS, GAME_SPRITES} from "../../model/storage/gameConstants.js";

export function drawBullets(context, state) {
    for (const bullet of Object.values(state.player.bullets)) {
        drawPLayerBullet(context, bullet);
    }

    for (const enemy of Object.values(state.enemies)) {
        for (const enemyBullet of Object.values(enemy.bullets)) {
            drawEnemyBullet(context, enemyBullet);
        }
    }
}

function drawPLayerBullet(context, bullet) {
    const sprite = GAME_SPRITES.BULLET_FLIES;

    context.save();
    context.translate(bullet.x, bullet.y);
    context.rotate(bullet.movementDirection + Math.PI / 2);
    context.drawImage(sprite, -GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
        GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
    context.restore();
}

function drawEnemyBullet(context, enemyBullet) {
    const sprite = GAME_SPRITES.ENEMY_BULLET_FLIES;

    context.save();
    context.translate(enemyBullet.x, enemyBullet.y);
    context.rotate(enemyBullet.movementDirection + Math.PI / 2);
    context.drawImage(sprite, -GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
        GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
    context.restore();
}