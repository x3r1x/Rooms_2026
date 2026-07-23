import {sendReadyState} from "../model/di/messages/clientMessages.js";

export function initLobbyListeners(socket) {
    let isReady = false;

    document.getElementById("readyButton").onclick = () => {
        isReady = !isReady;
        sendReadyState(socket, isReady);
        console.log("Send!");
    }
}