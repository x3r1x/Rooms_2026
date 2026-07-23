import {getSocket} from "../model/di/webSocket/server.js";

export function initRegistrationListeners(socket) {
    document.getElementById("nicknameButton").onclick = () => {
        const nicknameInput = document.getElementById("nicknameInput");
        socket = getSocket(nicknameInput.value);
    }
}