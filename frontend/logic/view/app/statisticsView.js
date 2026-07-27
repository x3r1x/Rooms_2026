export function updateStatisticView(statistic, gameNicknames, clientId) {
    const listElement = document.getElementById("players-list");

    statistic.sort((stat1, stat2) => statisticComparator(stat1, stat2))

    listElement.innerHTML = statistic.map(stat => {
        const isMe = stat.id === clientId;
        let hp = Math.round(stat.h);
        if (hp <= 0){
            hp = "dead"
        }
        return `
            <div class="player-stat-item ${isMe ? 'is-me' : ''}">
            <p>
                <span class="player-stat-name">${gameNicknames[stat.id]}</span>
                <span class="player-stat-hp">${hp.toString() === "dead" ? '' : 'HP:'} ${hp}</span>
            </p>
            <p>Kills: ${stat.k} / Deaths: ${stat.d}</p>
            </div>
        `;
    }).join('');
}

function statisticComparator(stat1, stat2) {
    if (stat1.k === stat2.k) {
        if (stat1.d === stat2.d) {
            return stat1.id.localeCompare(stat2.id);
        }

        return stat1.d - stat2.d;
    }

    return stat2.k - stat1.k;
}