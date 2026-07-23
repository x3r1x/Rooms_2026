import {APP_STATES} from "./appConstants.js";
import {
    showCountdownWindow,
    switchWindowFromLobbyToGame,
    switchWindowFromRegistrationToLobby,
    switchWindowToGameEnd
} from "../../view/app/windowSwitcher.js";
import {initLobbyListeners} from "../../controller/lobbyListeners.js";
import {loadGame, startGameLoop} from "../game/loadGame.js";
import {startGameState} from "../game/storage/gameState.js";
import {initGameEndListeners} from "../../controller/gameEndListeners.js";
import {fillResultWindow} from "../../view/app/gameEndView.js";

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

export function switchToCountdownAppState(countdown, lobbyState, gameNicknames) {
    if (appState !== APP_STATES.COUNTDOWN) {
        appState = APP_STATES.COUNTDOWN;

        lobbyState.players.forEach((player) => {
            gameNicknames[player.id] = player.n;
        })

        loadGame()
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