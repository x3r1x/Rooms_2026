import {APP_STATES} from "./appConstants.js";
import {
    showCountdownWindow,
    switchWindowFromLobbyToGame,
    switchWindowFromRegistrationToLobby,
    switchWindowToGameEnd,
    switchWindowToRegistration
} from "../../view/app/windowSwitcher.js";
import {initLobbyListeners} from "../../controller/lobbyListeners.js";
import {loadGame, startGameLoop} from "../game/loadGame.js";
import {finalStatistics, gameNicknames, gameState, startGameState} from "../game/storage/gameState.js";
import {initGameEndListeners} from "../../controller/gameEndListeners.js";
import {fillResultWindow} from "../../view/app/gameEndView.js";import {initRegistrationListeners} from "../../controller/registrationListeners.js";

export let appState = null
export let socket = null

export const lobbyState = {
    clientId: null,
    players: {},
    countdown: null
}

export function initAppState() {
    appState = APP_STATES.REGISTRATION;
}

export function switchToRegistrationAppState() {
    appState = APP_STATES.REGISTRATION;
    resetStates();
    initRegistrationListeners(socket);
    switchWindowToRegistration();
}

export function switchToWaitingAppState(socket, clientId) {
    if (appState !== APP_STATES.WAITING) {
        appState = APP_STATES.WAITING;
        initLobbyListeners(socket, clientId);
        switchWindowFromRegistrationToLobby();
        loadGame();
    }
}

export function switchToCountdownAppState(countdown, lobbyState, gameNicknames) {
    if (appState !== APP_STATES.COUNTDOWN) {
        appState = APP_STATES.COUNTDOWN;

        lobbyState.players.forEach((player) => {
            gameNicknames[player.id] = player.n;
        })

        showCountdownWindow(countdown);
    }
}

export function switchToOngoingGameState() {
    if (appState !== APP_STATES.GAME_ONGOING) {
        appState = APP_STATES.GAME_ONGOING;
        startGameState(performance.now());
        requestAnimationFrame(startGameLoop);
        switchWindowFromLobbyToGame();
    }
}

export function switchToEndedGameState(socket, clientId, result) {
    appState = APP_STATES.GAME_END;
    socket.close();
    initGameEndListeners();
    fillResultWindow(clientId, result);
    switchWindowToGameEnd();
}

export function resetStates() {
    appState = null;
    socket = null;
    lobbyState.clientId = null;
    lobbyState.players = {};
    lobbyState.countdown = null;
    gameState = null;
    gameNicknames = {};
    finalStatistics = null;
}