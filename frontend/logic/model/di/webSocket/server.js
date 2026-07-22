import {gameState, lobbyState} from "../../game/storage/gameState.js";
import {registerClient, sendGameInfo} from "../messages/clientMessages.js";
import {
    processWaitingMessage,
    processReadyMessage,
    processGameAssignment,
    processCountdownMessage
} from "../messages/serverMessages.js";
import {socket} from "../../../main.js";
import {APP_STATES} from "../../app/appConstants.js";
import {appState} from "../../app/appState.js";

export function getSocket(nickname) {
    const socket = new WebSocket("ws://84.201.159.214:8080/ws");

    socket.onopen = function (event) {
        console.log(`Открыто соединение - ${event.type}!`);
        registerClient(socket, nickname)
    }

    socket.onmessage = function (event) {
        parseMessage(event.data)
    }

    socket.onerror = function (error) {
        console.log(`Ошибка сокета - ${error.message}!`);
    }

    return socket;
}

export function parseMessage(message) {
    let parsedMessage = JSON.parse(message);

    switch (parsedMessage.s) {
        case APP_STATES.WAITING:
            appState = APP_STATES.WAITING;
            processWaitingMessage(parsedMessage, lobbyState);
            break;

        case APP_STATES.READY:
            appState = APP_STATES.COUNTDOWN;
            processReadyMessage(parsedMessage, lobbyState, gameState);
            break;

        case APP_STATES.COUNTDOWN:
            appState = APP_STATES.COUNTDOWN;
            processCountdownMessage(parsedMessage, lobbyState);
            break;

        case APP_STATES.GAME_ONGOING:
            appState = APP_STATES.GAME_ONGOING;
            sendGameInfo(socket, gameState);

            if (parsedMessage.type === "a") {
                processGameAssignment(parsedMessage, gameState);
            } else {
                console.log(`Packet loss: ${parsedMessage.type}!`);
            }
            break;
        default:
            console.log(`Непонятный state - ${parsedMessage.s}`);
    }
}