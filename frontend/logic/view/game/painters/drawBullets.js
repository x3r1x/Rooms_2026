import {GAME_CONSTANTS, GAME_SPRITES} from "../../../model/game/storage/gameConstants.js";
import {gameState} from "../../../model/game/storage/gameState.js";

export function drawBullets(context, playerId, bullets) {
    for (const bullet of Object.values(bullets)) {
        if (bullet.ownerId === playerId) {
            drawPlayerBullet(context, bullet);
        } else {
            drawEnemyBullet(context, bullet);
        }
    }
}

function drawPlayerBullet(context, bullet) {
    const sprite = GAME_SPRITES.PLAYER[`b${gameState.player.pc}`].img;

    context.save();
    context.translate(bullet.x, bullet.y);
    context.rotate(bullet.direction + Math.PI / 2);
    context.drawImage(sprite, -GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
        GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
    context.restore();
}

function drawEnemyBullet(context, enemyBullet) {
    //FIXME: игрок может быть вне комнаты, однако его пули всё ещё могут быть в комнате(тогда они не рисуются)
    if (gameState.enemies[enemyBullet.ownerId]) {
        const sprite = GAME_SPRITES.ENEMY[`b${gameState.enemies[enemyBullet.ownerId].pc}`].img;

        context.save();
        context.translate(enemyBullet.x, enemyBullet.y);
        context.rotate(enemyBullet.direction + Math.PI / 2);
        context.drawImage(sprite, -GAME_CONSTANTS.BULLET_WIDTH / 2, -GAME_CONSTANTS.BULLET_HEIGHT / 2,
            GAME_CONSTANTS.BULLET_WIDTH, GAME_CONSTANTS.BULLET_HEIGHT);
        context.restore();
    }
}