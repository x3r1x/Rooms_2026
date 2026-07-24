import {getSocket} from "../model/di/webSocket/server.js";

export function initRegistrationListeners(socket) {
    const button = document.getElementById("nicknameButton");
    button.onclick = () => {
        const nicknameInput = document.getElementById("nicknameInput");
        socket = getSocket(nicknameInput.value);
        button.disabled = true;
    }
}