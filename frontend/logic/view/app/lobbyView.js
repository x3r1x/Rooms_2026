const readyButton = document.getElementById("readyButton");
const playersConnected = document.getElementById("playersConnected");
const lobbyList = document.getElementById("lobbyList")

export function updateReadyText(isReady) {
    if (isReady) {
        readyButton.textContent = "NOT READY";
    }

    if (!isReady) {
        readyButton.textContent = "READY";
    }
}

export function updateLobbyView(ownId, playersInLobby) {
    playersConnected.textContent = playersInLobby.length;
    const readyStyle = `style="background-color: rgba(33, 255, 25, 0.5)"`;
    playersInLobby.sort((player1, player2) => player1.id.localeCompare(player2.id));

    lobbyList.innerHTML = playersInLobby.map(player => {
        const isPlayer = player.id === ownId;

        return `<p class="lobby-list-element" ${player.r === true ? readyStyle : ""}>${player.n} ${isPlayer ? "(You!)" : ""}</p>`;
    }).join('')
}