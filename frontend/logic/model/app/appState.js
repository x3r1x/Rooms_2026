import {APP_STATES} from "./appConstants.js";
import {
    showCountdownWindow,
    switchWindowFromLobbyToGame,
    switchWindowFromRegistrationToLobby
} from "../../view/app/windowSwitcher.js";
import {initLobbyListeners} from "../../controller/lobbyListeners.js";
import {loadGame, startGameLoop} from "../game/loadGame.js";
import {startGameState} from "../game/storage/gameState.js";

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

export function switchToWaitingAppState(socket) {
    if (appState !== APP_STATES.WAITING) {
        appState = APP_STATES.WAITING;
        initLobbyListeners(socket);
        switchWindowFromRegistrationToLobby();
    }
}

export function switchToCountdownAppState(countdown) {
    if (appState !== APP_STATES.COUNTDOWN) {
        appState = APP_STATES.COUNTDOWN;
        loadGame();
        showCountdownWindow(countdown);
    }
}

export function switchToOngoingGameState() {
    if (appState !== APP_STATES.GAME_ONGOING) {
        appState = APP_STATES.GAME_ONGOING;
        startGameState(performance.now())
        requestAnimationFrame(startGameLoop)
        switchWindowFromLobbyToGame();
    }
}