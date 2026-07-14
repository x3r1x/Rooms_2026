import {GAME_CONSTANTS, GAME_SPRITES} from "../../model/storage/gameConstants.js";

export function drawPlayers(context, state) {
    drawMainPlayer(context, state.player);

    state.enemies.forEach((enemy) => drawEnemy(context, enemy));
}

function drawMainPlayer(context, player) {
    const sprite = GAME_SPRITES.PLAYER_GOES;
    const size = GAME_CONSTANTS.PLAYER_VISUAL_SIZE;

    context.save();
    context.translate(player.x, player.y);
    context.rotate(player.direction);
    context.drawImage(sprite, -size / 2, -size / 2, size, size);
    context.restore();
}

function drawEnemy(context, enemy) {
    const sprite = GAME_SPRITES.ENEMY_GOES;
    const size = GAME_CONSTANTS.PLAYER_VISUAL_SIZE;

    context.save();
    context.translate(enemy.x, enemy.y);
    context.rotate(enemy.direction);
    context.drawImage(sprite, -size / 2, -size / 2, size, size);
    context.restore();
}