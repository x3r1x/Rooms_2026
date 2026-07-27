import {GAME_SPRITES} from "../../../model/game/storage/gameConstants.js";
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
    const type = `b${bullet.type}`;
    const sprite = GAME_SPRITES.PLAYER[`b${bullet.type}`].img;

    context.save();
    context.translate(bullet.x, bullet.y);
    context.rotate(bullet.direction);
    context.drawImage(sprite, -GAME_SPRITES.PLAYER[type].w / 2, -GAME_SPRITES.PLAYER[type].h / 2,
        GAME_SPRITES.PLAYER[type].w, GAME_SPRITES.PLAYER[type].h);
    context.restore();
}

function drawEnemyBullet(context, enemyBullet) {
    const type = `b${enemyBullet.type}`;
    const sprite = GAME_SPRITES.ENEMY[type].img;

    context.save();
    context.translate(enemyBullet.x, enemyBullet.y);
    context.rotate(enemyBullet.direction);
    context.drawImage(sprite, -GAME_SPRITES.ENEMY[type].w/ 2, -GAME_SPRITES.ENEMY[type].h / 2,
        GAME_SPRITES.ENEMY[type].w, GAME_SPRITES.ENEMY[type].h);
    context.restore();
}