import {currentState, lastState} from "../game/gameModel.js";
import {processInterpolation} from "../../controller/server/processInterpolation.js";
import {processAssignment} from "../../controller/server/processAssignment.js";

export function initSocket() {
    const socket = new WebSocket("ws://84.201.159.214:8080/ws");

    socket.onopen = function (event) {
        console.log(`Открыто соединение - ${event.type}!`);
    }

    socket.onmessage = function (event) {
        parseMessage(event.data, currentState, lastState);
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

export function parseMessage(message, state, previousState) {
    let parsedMessage = JSON.parse(message);

    if (parsedMessage.type === "interpolation" && "playerInterpolations" in parsedMessage) {
        processInterpolation(parsedMessage["playerInterpolations"], state, previousState);
    } else if (parsedMessage.type === "absolute" && "players" in parsedMessage) {
        processAssignment(parsedMessage["players"], state, previousState);
    } else {
        console.log(`Packet loss: ${parsedMessage.type}!`);
    }
}