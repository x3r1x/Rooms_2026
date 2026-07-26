export function updateStatisticView(gameState, gameNicknames) {
    const listElement = document.getElementById("players-list");

    let allPlayers = [gameState.player];

    for (const enemy of Object.values(gameState.enemies)) {
        allPlayers.push(enemy);
    }

    allPlayers.sort((p1, p2) => p1.id.localeCompare(p2.id));

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