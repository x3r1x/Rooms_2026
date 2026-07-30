import {finalStatistics, gameNicknames, gameState, setFinalStatistics} from "../../game/storage/gameState.js";
import {registerClient, sendGameInfo} from "./messages/clientMessages.js";
import {
    processCountdownMessage,
    processGameAssignment,
    processReadyMessage,
    processWaitingMessage
} from "./messages/serverMessages.js";
import {APP_STATES} from "../../app/appConstants.js";
import {
    lobbyState,
    switchToCountdownAppState,
    switchToEndedGameState,
    switchToOngoingGameState,
    switchToWaitingAppState
} from "../../app/appState.js";

export function getSocket(nickname) {
    const socket = new WebSocket("ws://84.201.159.214:8080/ws");
    // const socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = function (event) {
        console.log(`Открыто соединение - ${event.type}!`);
        registerClient(socket, nickname);
    }

    socket.onmessage = function (event) {
        parseMessage(socket, event.data);
    }

    socket.onerror = function (error) {
        console.log(`Ошибка сокета - ${error.message}!`);
        socket.close();
    }

    return socket;
}

export function parseMessage(socket, message) {
    let parsedMessage = JSON.parse(message);

    switch (parsedMessage.s) {
        case APP_STATES.WAITING:
            processWaitingMessage(parsedMessage, lobbyState);
            switchToWaitingAppState(socket, lobbyState.clientId);
            break;

        case APP_STATES.READY:
            switchToCountdownAppState(parsedMessage.c, lobbyState, gameNicknames)
            processReadyMessage(parsedMessage, lobbyState, gameState);
            break;

        case APP_STATES.COUNTDOWN:
            switchToCountdownAppState(parsedMessage.c, lobbyState, gameNicknames)
            processCountdownMessage(parsedMessage, lobbyState);
            break;

        case APP_STATES.GAME_ONGOING:
            if (parsedMessage.k.length !== 0) {
                console.log("There!");
            }
            switchToOngoingGameState();
            sendGameInfo(socket, gameState);

            processGameAssignment(parsedMessage, gameState, gameNicknames);
            break;

        case APP_STATES.GAME_END:
            setFinalStatistics(parsedMessage.r);
            switchToEndedGameState(socket, gameState.player.id, finalStatistics);
            break;

        default:
            console.log(`Непонятный state - ${parsedMessage.s}`);
    }
}