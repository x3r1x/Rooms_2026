import {layersForRoom, room} from "./gameState.js";

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
        })
        layersForRoom.tilesInfo[tile.id] = {
            id: tile.id,
            blocksBullet: property.blocksBullet,
            blocksPlayer: property.blocksPlayer,
            hitboxes: property.hitbox,
        };
    });
}

export function getMapCollision(){
    room.collision = assemblyObject(room.walls, room.object);
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
    let layersAnswer = [...layers[0].data];

    for (let i = 1; i < layers.length; i++) {
        let tempObj = assemblyObject({ data: layersAnswer }, layers[i]);
        layersAnswer = tempObj.data;
    }

    return {name: "merge", data: layersAnswer};
}

export function getFlapList(arrayExit) {
    const flaps = {top: true, left: true, down: true, right: true};
    for (let exit of arrayExit) {
        const exitLower = exit.name.toLowerCase();
        if (exitLower.includes('top')) {
            flaps.top = false;
        }
        if (exitLower.includes('left')) {
            flaps.left = false;
        }
        if (exitLower.includes('down')) {
            flaps.down = false;
        }
        if (exitLower.includes('right')) {
            flaps.right = false;
        }
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

export function assemblyRoom() {
    const arrayExit = getRandomArray(layersForRoom.exit);

    room.exit = checkExit(arrayExit);
    room.flap = getFlapList(arrayExit);
    const arrayFlap = getFlap(room.flap);

    const layersFlap = mergeLayers(arrayFlap);
    const layersExit = mergeLayers(arrayExit);

    room.exits = assemblyObject(layersForRoom.floor[0], layersExit);
    room.walls = assemblyObject(layersForRoom.walls[0], layersFlap);

    let randomIndexObjects = Math.floor(Math.random() * (layersForRoom.objects.length));
    room.object = layersForRoom.objects[randomIndexObjects];

    room.floors = layersForRoom.floor[0];

    getMapCollision();
}

