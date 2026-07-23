import {changeCountdownTimer} from "./countdownView.js";

const nicknameWrapper = document.getElementById("nicknameWrapper");
const lobbyWrapper = document.getElementById("lobbyWrapper");
const mainGameWrapper = document.getElementById("mainGameWrapper");
const gameOverWrapper = document.getElementById("gameOverWrapper");
const countdownWrapper = document.getElementById("countdownWrapper");

export function switchWindowFromRegistrationToLobby() {
    nicknameWrapper.style.display = "none";
    lobbyWrapper.style.display = "flex";
}

export function showCountdownWindow(timeRemaining) {
    countdownWrapper.style.display = "flex";

    changeCountdownTimer(timeRemaining);
}

export function switchWindowFromLobbyToGame() {
    lobbyWrapper.style.display = "none";
    countdownWrapper.style.display = "none"
    mainGameWrapper.style.display = "flex";
}

export function switchWindowToGameEnd() {
    mainGameWrapper.style.display = "none";
    gameOverWrapper.style.display = "flex";
}