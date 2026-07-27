import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";
import {drawPlayers} from "./drawPlayers.js";
import {updateAndDrawEffects} from "../effects/effectsManager.js";

export function drawGame(canvas, context, state) {
    drawBackground(canvas, context);
    drawBullets(context, state.player.id, state.bullets);
    drawPlayers(context, state.player, state.enemies);

    updateAndDrawEffects(context);

    context.fill();
}