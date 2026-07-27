const readyButton = document.getElementById("readyButton");
const playersConnected = document.getElementById("playersConnected");
const lobbyList = document.getElementById("lobbyList")

export let selectedWeaponClass = "g";
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

export function updatePlayerClass(){
    const weapons = [
        { name: "ПИСТОЛЕТ", img: "/frontend/assets/images/g.png", code: "g"},
        { name: "ДРОБОВИК", img: "/frontend/assets/images/r.png", code: "r" },
        { name: "СОМ", img: "/frontend/assets/images/s.png", code: "s"},
    ];

    let currentIndex = 0;

    const btn = document.getElementById('changeWeaponBtn');
    const nameSpan = document.getElementById('weaponName');
    const imageImg = document.getElementById('weaponImage');

    btn.addEventListener('click', () => {
        currentIndex = (currentIndex + 1) % weapons.length;
        const currentWeapon = weapons[currentIndex];
        nameSpan.textContent = currentWeapon.name;
        selectedWeaponClass = currentWeapon.code;
        imageImg.src = currentWeapon.img;
    });
}

