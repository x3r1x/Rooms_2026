import {processAssignment} from "./processAssignment.js";
import {currentState} from "../../storage/states.js";

export function getSocket() {
    const socket = new WebSocket("ws://84.201.159.214:8080/ws");

    socket.onopen = function (event) {
        console.log(`Открыто соединение - ${event.type}!`);
        sendMessage(socket, currentState)
    }

    socket.onmessage = function (event) {
        sendMessage(socket, currentState);
        parseMessage(event.data, currentState);

        currentState.didShoot = false;
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
    if (parsedMessage.type === "a") {
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
        "s": state.player.didShoot
    }

    socket.send(JSON.stringify(message));
}