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

/*export function parserTileInfo(dataJson) {
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
}*/

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
        return { name: "empty", data: [] };
    }
    let layersAnswer = [...layers[0].data];

    for (let i = 1; i < layers.length; i++) {
        let tempObj = assemblyObject({ data: layersAnswer }, layers[i]);
        layersAnswer = tempObj.data;
    }

    return {name: "merge", data: layersAnswer};
}

function checkExit(layers) {
    const exits = {top: false, left: false, down: false, right: false};
    for (let layer of layers) {
        const layerLower = layer.name.toLowerCase();
        if (layerLower.includes('top')) exits.top = true;
        if (layerLower.includes('left')) exits.left = true;
        if (layerLower.includes('down')) exits.down = true;
        if (layerLower.includes('right')) exits.right = true;
    }
    return exits;
}

export function getFlapList(arrayExit) {
    const flaps = {top: true, left: true, down: true, right: true};
    for (let exit of arrayExit) {
        const exitLower = exit.name.toLowerCase();
        if (exitLower.includes('top')) flaps.top = false;
        if (exitLower.includes('left')) flaps.left = false;
        if (exitLower.includes('down')) flaps.down = false;
        if (exitLower.includes('right')) flaps.right = false;
    }
    return flaps;
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

    const baseFloor = layersForRoom.floor[0] || { name: "empty", data: [] };
    const baseWalls = layersForRoom.walls[0] || { name: "empty", data: [] };

    targetRoom.exits = assemblyObject(baseFloor, layersExit);
    targetRoom.walls = assemblyObject(baseWalls, layersFlap);

    const matchingObjects = layersForRoom.objects.filter(obj =>
        obj.name.toLowerCase().includes(targetRoom.type.toLowerCase())
    );

    if (matchingObjects.length > 0) {
        let randomIndexObjects = Math.floor(Math.random() * matchingObjects.length);
        targetRoom.object = matchingObjects[randomIndexObjects];
    } else {
        targetRoom.object = { name: "empty", data: [] };
    }

    targetRoom.floors = baseFloor;
    console.log("targetRoom", targetRoom);
    getMapCollision(targetRoom);
}