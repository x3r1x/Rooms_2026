import {processInterpolation} from "./processInterpolation.js";
import {processAssignment} from "./processAssignment.js";
import {currentState, previousState} from "../../storage/states.js";

export function initSocket() {
    const socket = new WebSocket("ws://84.201.159.214:8080/ws");

    socket.onopen = function (event) {
        console.log(`Открыто соединение - ${event.type}!`);
    }

    socket.onmessage = function (event) {
        sendMessage(socket, currentState, previousState);
        parseMessage(event.data, previousState);
        processCurrentState(currentState, previousState);
    }

    socket.onerror = function (error) {
        console.log(`Ошибка сокета - ${error.message}!`);
    }

    return socket;
}

export function parseMessage(message, state) {
    let parsedMessage = JSON.parse(message);

    if (parsedMessage.type === "interpolation" && "playerInterpolations" in parsedMessage) {
        processInterpolation(parsedMessage["playerInterpolations"], state);
    } else if (parsedMessage.type === "absolute" && "players" in parsedMessage) {
        processAssignment(parsedMessage["players"], state);
    } else {
        console.log(`Packet loss: ${parsedMessage.type}!`);
    }
}

function sendMessage(socket, state, previousState) {
    const message = {
        "playerInterpolation": {
            "direction": state.movementDirection,
            "deltaDirection": state.player.direction - previousState.player.direction,
            "id": state.player.id,
            "newBulletsDirection": state.newBulletsDirection
        }
    }

    socket.send(JSON.stringify(message));
}

function processCurrentState(currentState, previousState) {
    currentState.player = previousState.player;
    currentState.enemies = previousState.enemies;

    currentState.movementDirection = {
        x: 0,
        y: 0
    }
    currentState.newBulletsDirection = [];
}