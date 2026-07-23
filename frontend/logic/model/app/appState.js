import {APP_STATES} from "./appConstants.js";
import {switchWindowFromRegistrationToLobby} from "../../view/app/windowSwitcher.js";
import {initLobbyListeners} from "../../controller/lobbyListeners.js";

export let appState = null
export let socket = null

export function initAppState() {
    appState = APP_STATES.REGISTRATION;
}

export function switchToWaitingAppState(socket) {
    appState = APP_STATES.WAITING;
    initLobbyListeners(socket);
    switchWindowFromRegistrationToLobby();
}

export function switchToGameOngoingAppState() {
    appState = APP_STATES.GAME_ONGOING;
}