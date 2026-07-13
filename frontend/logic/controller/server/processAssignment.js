import {getPlayerFromModelById} from "../engine/player.js";

export function processAssignment(players, state, previousState) {
    for (const [id, player] of Object.entries(players)) {
        let playerInModel = getPlayerFromModelById(previousState, id);

        if (playerInModel === null) {
            console.log(`Unknown id: ${id}!`);
        } else {
            assignInfoToModel(playerInModel, player);
        }
    }
}

function assignInfoToModel(playerInModel, info) {
    playerInModel.x = info.x;
    playerInModel.y = info.y;
    playerInModel.direction = info.direction;
    playerInModel.bullets = info.bullets;
}