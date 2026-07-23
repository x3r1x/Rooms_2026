export function updateStatisticView(gameState) {
    const listElement = document.getElementById("players-list");

    const allPlayers = [gameState.player, ...gameState.enemies];

    listElement.innerHTML = allPlayers.map(p => {
        const isMe = p.id === gameState.player.id;
        let hp = p.hp;
        if (hp <= 0){
            hp = "dead"
        }
        return `
            <p class="player-stat-item ${isMe ? 'is-me' : ''}">
                <span class="player-stat-name">${p.id}</span>
                <span class="player-stat-hp">${hp==="dead" ? '' : 'HP:'} ${hp}</span>
            </p>
        `;
    }).join('');
}