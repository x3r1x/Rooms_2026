import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";
import {GAME_CONSTANTS, GAME_SPRITES} from "../gameConstants.js";

export function drawGame(canvas, context, state) {
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawPlayer(canvas, context, state);

    context.fill();
}

function drawPlayer(canvas, context, state) {
    const player = state.player;
    const sprite = GAME_SPRITES.PLAYER_GOES;
    const size = GAME_CONSTANTS.PLAYER_VISUAL_SIZE;
    context.drawImage(sprite, player.x, player.y, size, size);
}