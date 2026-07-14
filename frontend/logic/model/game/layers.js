import {layersForRoom, room} from "./gameModel.js";

export function parserMapData(dataJson) {
    layersForRoom.width = dataJson.width;
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

function getRandomArray(array) {
    const lenArray = Math.floor(Math.random() * array.length) + 1;
    const result = [];

    for (let i = 0; i < lenArray; i++) {
        const randomIndex = Math.floor(Math.random() * array.length);
        result[i] = array[randomIndex];
    }

    return result;
}

function assemblyObject(object1, object2) {
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

function mergeLayers(layers) {
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
        if (layerLower.includes('top')) {
            exits.top = true;
        }
        if (layerLower.includes('left')) {
            exits.left = true;
        }
        if (layerLower.includes('down')) {
            exits.down = true;
        }
        if (layerLower.includes('right')) {
            exits.right = true;
        }
    }
    return exits;
}

function getFlapList(arrayExit) {
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

function getFlap(flapList) {
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

}

