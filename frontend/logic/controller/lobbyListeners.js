import {sendReadyState} from "../model/di/messages/clientMessages.js";
import {updateReadyText} from "../view/app/lobbyView.js";

let isReady = false;

export function initLobbyListeners(socket) {
    document.getElementById("readyButton").onclick = () => {
        isReady = !isReady;
        updateReadyText(isReady);
        sendReadyState(socket, isReady);
    }
}