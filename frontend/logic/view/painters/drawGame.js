import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";
import {GAME_CONSTANTS, GAME_SPRITES} from "../../model/gameConstants.js";

export function drawGame(canvas, context, state) {
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawPlayers(canvas, context, state);

    context.fill();
}

function drawPlayers(canvas, context, state) {
    const sprite = GAME_SPRITES.PLAYER_GOES;
    const size = GAME_CONSTANTS.PLAYER_VISUAL_SIZE;
    context.save();
    context.translate(state.player.x, state.player.y);
    context.rotate(state.player.direction);
    context.drawImage(sprite, -size/2, -size/2, size, size);
    context.restore();

    context.fillStyle = "#ff0000";

    state.enemies.forEach(function (enemy) {
        context.fillRect(enemy.x, enemy.y, GAME_CONSTANTS.SQUARE_VISUAL_SIZE, GAME_CONSTANTS.SQUARE_VISUAL_SIZE);
    })
}