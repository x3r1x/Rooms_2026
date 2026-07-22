export function processAssignment(parsedMessage, state) {
    state.enemies = [];
    parsedMessage.p.forEach((player) => processPlayer(player, state));

    state.bullets = [];
    parsedMessage.b.forEach((bullet) => processBullet(bullet, state));

    updatePlayersStatistic(state);
}

function processPlayer(player, state) {
    const newPlayerInModel = {
        x: player.x,
        y: player.y,
        direction: player.a,
        movementDirection: {
            x: player.mx,
            y: player.my
        },
        id: player.id,
        hp: player.h,
        rebornTime: player.rt
    }

    if (newPlayerInModel.id === state.player.id) {
        newPlayerInModel.mousePosition = state.player.mousePosition;
        state.player = newPlayerInModel;
    } else {
        state.enemies.push(newPlayerInModel);
    }
}

function processBullet(bullet, state) {
    state.bullets.push({
        x: bullet.x,
        y: bullet.y,
        direction: bullet.a,
        ownerId: bullet.oId
    })
}

function updatePlayersStatistic(state) {
    const listElement = document.getElementById("players-list");


    const allPlayers = [state.player, ...state.enemies];

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