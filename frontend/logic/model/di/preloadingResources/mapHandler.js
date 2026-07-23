import {layersForRoom, map} from "../../game/storage/gameState.js";
import {assemblyObject, getFlap, getFlapList, mergeLayers} from "../../game/storage/layers.js";

const dataJson = {
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
}


export function parseMap() {
    if (!dataJson || !('m' in dataJson)) {
        return;
    }
    const roomsMap = dataJson.m;

    for (const uuid in roomsMap) {
        const currentRoom = roomsMap[uuid];
        map[uuid] = {
            id: currentRoom.id,
            type: currentRoom.bT,
            doors: {
                top: currentRoom.eT || null,
                left: currentRoom.eL || null,
                bottom: currentRoom.eB || null,
                right: currentRoom.eR || null
            },
            walls: [],
            object: [],
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
    for (const uuid in map) {
        const currentRoom = map[uuid];

        const arrayFlap = getFlap(currentRoom.flap);

        let assembledWalls = layersForRoom.walls[0];
        console.log(assembledWalls);
        if (arrayFlap.length > 0) {
            const layersFlap = mergeLayers(arrayFlap);
            assembledWalls = assemblyObject(layersForRoom.walls[0], layersFlap);
        }
        console.log(layersForRoom.walls);

        const targetObjectName = `barrier${currentRoom.type}`;
        const objectLayer = layersForRoom.objects.find(obj => obj.name.toLowerCase() === targetObjectName.toLowerCase());

        currentRoom.object = objectLayer ? objectLayer.data : [];

        if (objectLayer) {
            const collisionResult = assemblyObject(assembledWalls, objectLayer);
            currentRoom.collision = collisionResult.data;
        } else {
            currentRoom.collision = assembledWalls.data;
        }
    }
}