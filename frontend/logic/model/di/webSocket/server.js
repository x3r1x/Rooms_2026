import {processAssignment} from "./processAssignment.js";
import {currentState, previousState} from "../../storage/states.js";

export function getSocket() {
    const socket = new WebSocket("ws://84.201.159.214:8080/ws");

    socket.onopen = function (event) {
        console.log(`Открыто соединение - ${event.type}!`);
    }

    socket.onmessage = function (event) {
        sendMessage(socket, currentState);
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

    // if (parsedMessage.type === "interpolation" && "playerInterpolations" in parsedMessage) {
    //     processInterpolation(parsedMessage["playerInterpolations"], state);
    if (parsedMessage.type === "a" && "players" in parsedMessage) {
        processAssignment(parsedMessage, state);
    } else {
        console.log(`Packet loss: ${parsedMessage.type}!`);
    }
}

function sendMessage(socket, state) {
    const message = {
        "id": state.player.id,
        "a": state.player.direction,
        "mx": state.player.movementDirection.x,
        "my": state.player.movementDirection.y,
        "s": state.player.didShot
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
    currentState.didShot = false;
}