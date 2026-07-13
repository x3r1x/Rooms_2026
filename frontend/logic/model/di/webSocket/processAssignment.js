import {getPlayerFromModelById} from "../../engine/players.js";

export function processAssignment(players, state) {
    for (const [id, player] of Object.entries(players)) {
        let playerInModel = getPlayerFromModelById(state, id);

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
    playerInModel.direction = info.movementDirection;
    playerInModel.bullets = info.bullets;
}