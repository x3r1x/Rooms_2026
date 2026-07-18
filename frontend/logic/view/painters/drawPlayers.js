import {GAME_CONSTANTS, GAME_SPRITES} from "../../model/storage/gameConstants.js";
import {keys} from '../../controller/listeners.js'
export function drawPlayers(context, mainPlayer, enemies) {
    drawMainPlayer(context, mainPlayer);

    enemies.forEach((enemy) => drawEnemy(context, enemy));
}

function drawMainPlayer(context, player) {
    const spriteSheet = GAME_SPRITES.PLAYER_GOES;
    const frameWidth = spriteSheet.width / 6;
    const frameHeight = spriteSheet.height;
    if (player.movementDirection.x !== 0 || player.movementDirection.y !== 0) {
        player.spriteIndex = Math.floor(Date.now() / 200) % 6;
    } else {
        player.spriteIndex = 0;
    }
    context.save();
    context.translate(player.x, player.y);
    context.rotate(player.direction);
    context.drawImage(
        spriteSheet,
        player.spriteIndex * frameWidth, 0, frameWidth, frameHeight,
        -GAME_CONSTANTS.PLAYER_VISUAL_WIDTH/2, -GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT/2, GAME_CONSTANTS.PLAYER_VISUAL_WIDTH, GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT
    );
    context.restore();
}

function drawEnemy(context, enemy) {
    const spriteSheet = GAME_SPRITES.ENEMY_GOES;
    const frameWidth = spriteSheet.width / 6;
    const frameHeight = spriteSheet.height;
    if (enemy.movementDirection.x !== 0 || enemy.movementDirection.y !== 0) {
        enemy.spriteIndex = Math.floor(Date.now() / 200) % 6;
    } else {
        enemy.spriteIndex = 0;
    }
    context.save();
    context.translate(enemy.x, enemy.y);
    context.rotate(enemy.direction);
    context.drawImage(
        spriteSheet,
        enemy.spriteIndex * frameWidth, 0, frameWidth, frameHeight,
        -GAME_CONSTANTS.PLAYER_VISUAL_WIDTH/2, -GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT/2, GAME_CONSTANTS.PLAYER_VISUAL_WIDTH, GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT
    );
    context.restore();
}