const nicknameWrapper = document.getElementById("nicknameWrapper");
const lobbyWrapper = document.getElementById("lobbyWrapper");
const mainGameWrapper = document.getElementById("mainGameWrapper");
const gameOverWrapper = document.getElementById("gameOverWrapper");

export function switchWindowFromRegistrationToLobby() {
    nicknameWrapper.style.display = "none";
    lobbyWrapper.style.display = "flex";
}

export function switchWindowFromLobbyToGame() {
    lobbyWrapper.style.display = "none";
    mainGameWrapper.style.display = "flex";
}

export function switchWindowToGameEnd() {
    mainGameWrapper.style.display = "none";
    gameOverWrapper.style.display = "flex";
}