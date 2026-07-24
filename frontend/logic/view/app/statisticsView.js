export function updateStatisticView(gameState, gameNicknames) {
    const listElement = document.getElementById("players-list");

    const allPlayers = [gameState.player, ...gameState.enemies];
    allPlayers.sort((p1, p2) => gameNicknames[p1.id].localeCompare(gameNicknames[p2.id]));

    listElement.innerHTML = allPlayers.map(p => {
        const isMe = p.id === gameState.player.id;
        let hp = Math.round(p.hp);
        if (hp <= 0){
            hp = "dead"
        }
        return `
            <p class="player-stat-item ${isMe ? 'is-me' : ''}">
                <span class="player-stat-name">${gameNicknames[p.id]}</span>
                <span class="player-stat-hp">${hp.toString() === "dead" ? '' : 'HP:'} ${hp}</span>
            </p>
        `;
    }).join('');
}