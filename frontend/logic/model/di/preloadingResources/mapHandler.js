import {gameMap} from "../../game/storage/gameState.js";
import {assemblyRoom} from "../../game/storage/layers.js";

export function parseMap(dataJson) {
    if (!dataJson || !('m' in dataJson)) {
        return;
    }
    const roomsMap = dataJson.m;
    for (const uuid in roomsMap) {

        const currentRoom = roomsMap[uuid];
        gameMap[uuid] = {
            id: currentRoom.id,
            type: `barrier${currentRoom.bT}`,
            doors: {
                top: currentRoom.eT || null,
                left: currentRoom.eL || null,
                bottom: currentRoom.eB || null,
                right: currentRoom.eR || null
            },
            exits: [],
            floors: [],
            walls: [],
            object: [],
            exit: {
                top: Boolean(currentRoom.eT),
                left: Boolean(currentRoom.eL),
                down: Boolean(currentRoom.eB),
                right: Boolean(currentRoom.eR)
            },
            flap: {
                top: !currentRoom.eT,
                left: !currentRoom.eL,
                down: !currentRoom.eB,
                right: !currentRoom.eR
            },
            collision: []
        };
    }
    updateMap();
}

function updateMap() {
    for (const uuid in gameMap) {
        assemblyRoom(gameMap[uuid]);
    }
    console.log(gameMap);
}