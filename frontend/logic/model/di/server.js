import {updatePlayers} from "../../controller/engine/player.js";
import {lastState} from "../gameLogic/gameModel.js";

export function initSocket() {
    const socket = new WebSocket("ws://84.201.159.214:8080/ws");

    socket.onopen = function (event) {
        console.log(`Открыто соединение - ${event.type}!`);
    }

    socket.onmessage = function (event) {
        parseMessage(event.data, lastState);
    }

    socket.onerror = function (error) {
        console.log(`Ошибка сокета - ${error.message}!`);
    }

    return socket;
}

export function sendMessage(socket, playerInfo) {
    const message = {
        "player": playerInfo
    }

    socket.send(JSON.stringify(message));
}

function parseMessage(message, state) {
    let parsedMessage = JSON.parse(message);

    updatePlayers(parsedMessage["players"], state)
}