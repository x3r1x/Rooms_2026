
    const allPlayers = [state.player, ...state.enemies];

    allPlayers.sort((a, b) => a.id.localeCompare(b.id));

    listElement.innerHTML = allPlayers.map(p => {
        const isMe = p.id === state.player.id;
        let hp = p.hp;
        if (hp <= 0){
            hp = "kill"
        }
           return `
                <p class="player-stat-item ${isMe ? 'is-me' : ''}">
                    <span class="player-stat-name">${p.id}</span>
                    <span class="player-stat-hp">${hp=="kill" ? '' : 'HP:'} ${hp}</span>
                </p>
            `;
        }).join('');
}