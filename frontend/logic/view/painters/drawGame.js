import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";
import {GAME_CONSTANTS} from "../../model/gameLogic/gameConstants.js";

export function drawGame(canvas, context, state) {
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawPlayers(canvas, context, state);

    context.fill();
}

function drawPlayers(canvas, context, state) {
    context.fillStyle = "#FFFFFF";
    context.fillRect(state.square.x, state.square.y, GAME_CONSTANTS.SQUARE_VISUAL_SIZE, GAME_CONSTANTS.SQUARE_VISUAL_SIZE);

    context.fillStyle = "#ff0000";

    state.enemies.forEach(function (enemy) {
        context.fillRect(enemy.x, enemy.y, GAME_CONSTANTS.SQUARE_VISUAL_SIZE, GAME_CONSTANTS.SQUARE_VISUAL_SIZE);
    })
}