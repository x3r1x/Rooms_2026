import {gameMap} from "../../game/storage/gameState.js";
import {assemblyRoom} from "../../game/storage/layers.js";

/*const dataJson = {
    "s": "r",
    "c": 5.0,
    "m": {
        "UUID2": {
            "eT": "",
            "eL": "UUID1",
            "eB": "",
            "eR": "",
            "id": "UUID2",
            "bT": 2
        },
        "UUID1": {
            "eT": "",
            "eL": "",
            "eB": "",
            "eR": "UUID2",
            "id": "UUID1",
            "bT": 5
        },
        "UUID3": {
            "eT": "",
            "eL": "",
            "eB": "",
            "eR": "UUID2",
            "id": "UUID3",
            "bT": 5
        }
    }
}*/


export function parseMap(dataJson) {
    if (!dataJson || !('m' in dataJson)) {
        return;
    }
    const roomsMap = dataJson.m;
    let count = 0;
    for (const uuid in roomsMap) {

        const currentRoom = roomsMap[uuid];
        gameMap[count] = {
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

    count ++;
    }
    updateMap();
}

function updateMap() {
    for (const uuid in gameMap) {
        assemblyRoom(gameMap[uuid]);
    }
    console.log(gameMap);
}