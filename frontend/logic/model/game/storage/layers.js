import {layersForRoom} from "./gameState.js";

export function parseMapData(dataJson) {
    layersForRoom.width = dataJson.width;
    layersForRoom.height = dataJson.height;
    dataJson.layers.forEach(layer => {
        if (layer.name === 'baseRoom')
            layersForRoom.floor.push({name: layer.name, data: layer.data});
        else if (layer.name === 'baseWalls')
            layersForRoom.walls.push({name: layer.name, data: layer.data});
        else if (layer.name.includes('exit'))
            layersForRoom.exit.push({name: layer.name, data: layer.data});
        else if (layer.name.includes('flap'))
            layersForRoom.flap.push({name: layer.name, data: layer.data});
        else if (layer.name.includes('barrier'))
            layersForRoom.objects.push({name: layer.name, data: layer.data});
    });
}

export function parseTileInfo(dataJson) {
    layersForRoom.tileSize = dataJson.tileheight;
    const tileArray = dataJson.tiles;
    tileArray.forEach(tile => {
        const property = {};
        tile.properties.forEach(p => {
            property[p.name] = JSON.parse(p.value);
        });
        const tileInfo = {
            id: tile.id,
            blocksBullet: property.blocksBullet,
            blocksPlayer: property.blocksPlayer,
            hitboxes: property.hitbox,
        };
        layersForRoom.tilesInfo[tile.id] = tileInfo;
    });
}

export function getMapCollision(targetRoom) {
    targetRoom.collision = assemblyObject(targetRoom.walls, targetRoom.object);
}

function getExit(exitList) {
    const exits = [];
    for (let exit in exitList) {
        if (exitList[exit] === true) {
            const found = layersForRoom.exit.find(f => f.name.toLowerCase().includes(exit));
            if (found) {
                exits.push({
                    name: found.name,
                    data: found.data
                });
            }
        }
    }
    return exits;
}

export function assemblyObject(object1, object2) {
    const resultData = [...object1.data];

    for (let i = 0; i < resultData.length; i++) {
        if (object2.data[i] !== 0) {
            resultData[i] = object2.data[i];
        }
    }

    return {
        name: "assembly",
        data: resultData
    };
}

export function mergeLayers(layers) {
    if (!layers || layers.length === 0 || !layers[0]?.data) {
        return {name: "empty", data: []};
    }
    let layersAnswer = [...layers[0].data];

    for (let i = 1; i < layers.length; i++) {
        let tempObj = assemblyObject({data: layersAnswer}, layers[i]);
        layersAnswer = tempObj.data;
    }

    return {name: "merge", data: layersAnswer};
}

export function getFlap(flapList) {
    const flaps = [];
    for (let flap in flapList) {
        if (flapList[flap] === true) {
            const found = layersForRoom.flap.find(f => f.name.toLowerCase().includes(flap));
            if (found) {
                flaps.push({
                    name: found.name,
                    data: found.data
                });
            }
        }
    }
    return flaps;
}

export function assemblyRoom(targetRoom) {
    if (!targetRoom || !targetRoom.flap || !targetRoom.exit) {
        return;
    }

    const arrayFlap = getFlap(targetRoom.flap);
    const layersFlap = mergeLayers(arrayFlap);

    const arrayExit = getExit(targetRoom.exit);
    const layersExit = mergeLayers(arrayExit);

    const baseFloor = layersForRoom.floor[0];
    const baseWalls = layersForRoom.walls[0];

    targetRoom.exits = assemblyObject(baseFloor, layersExit);
    targetRoom.walls = assemblyObject(baseWalls, layersFlap);

    const matchingObject = layersForRoom.objects.find(obj =>
        obj.name.toLowerCase().includes(targetRoom.type.toLowerCase())
    );

    if (matchingObject) {
        targetRoom.object = matchingObject;
    } else {
        targetRoom.object = { name: "empty", data: [] };
    }

    targetRoom.floors = baseFloor;
    getMapCollision(targetRoom);
}