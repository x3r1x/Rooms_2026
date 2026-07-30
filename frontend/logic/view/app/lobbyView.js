const readyButton = document.getElementById("readyButton");
const playersConnected = document.getElementById("playersConnected");
const lobbyList = document.getElementById("lobbyList")
const btn = document.getElementById('changeWeaponBtn');
export let selectedWeaponClass = "g";

export function updateReadyText(isReady) {
    if (isReady) {
        readyButton.textContent = "NOT READY";
        btn.disabled = true;
    }
    if (!isReady) {
        readyButton.textContent = "READY";
        btn.disabled = false;
    }
}

export function updateLobbyView(ownId, playersInLobby) {
    playersConnected.textContent = playersInLobby.length;
    const readyStyle = `style="background-color: rgba(33, 255, 25, 0.5)"`;
    playersInLobby.sort((player1, player2) => player1.id.localeCompare(player2.id));

    lobbyList.innerHTML = playersInLobby.map(player => {
        const isPlayer = player.id === ownId;

        return `
            <p class="lobby-list-element ${isPlayer ? 'is-me' : ''}" ${player.r === true ? readyStyle : ""}>
                ${player.n} ${isPlayer ? "(You!)" : ""}
            </p>
        `;
    }).join('')
}

export function updatePlayerClass() {
    const weapons = [
        {
            name: "ПИСТОЛЕТ",
            img: "/assets/images/g.png",
            code: "g",
            description: `
                <span class="weapon-title">ПИСТОЛЕТ</span><br>
                <span class="stat">Урон:</span> Средний<br>
                <span class="stat">Скорость пули:</span> Высокая<br>
                <span class="stat">Скорость игрока:</span> Максимальная<br>
                <span class="stat">Особенность:</span> Бесконечный боезапас и абсолютная надежность. Идеален для динамичных маневров.
            `
        },
        {
            name: "ДРОБОВИК",
            img: "/assets/images/r.png",
            code: "r",
            description: `
                <span class="weapon-title">ДРОБОЛОБИК</span><br>
                <span class="stat">Урон:</span> Огромный (вблизи)<br>
                <span class="stat">Скорость пули:</span> Дробь летит быстро, но быстро гаснет<br>
                <span class="stat">Скорость игрока:</span> Средняя<br>
                <span class="stat">Особенность:</span> Стреляет веером. Сносит толпы врагов мощным хвостовым ударом на короткой дистанции.
            `
        },
        {
            name: "СОМ",
            img: "/assets/images/s.png",
            code: "s",
            description: `
                <span class="weapon-title">СОМ</span><br>
                <span class="stat">Урон:</span> Колоссальный<br>
                <span class="stat">Скорость пули:</span> Медленная<br>
                <span class="stat">Скорость игрока:</span> Пониженная<br>
                <span class="stat">Особенность:</span> Наносит урон по площади. Оставляет за собой выжженную зону и панику на всей карте.
            `
        },
    ];

    let currentIndex = 0;

    const imageImg = document.getElementById('weaponImage');
    const descriptionP = document.getElementById('weaponDescription');

    function updateWeapon() {
        const currentWeapon = weapons[currentIndex];

        imageImg.src = currentWeapon.img;
        selectedWeaponClass = currentWeapon.code;
        descriptionP.innerHTML = currentWeapon.description;
    }

    updateWeapon();

    btn.addEventListener('click', () => {
        currentIndex = (currentIndex + 1) % weapons.length;
        updateWeapon();
    });
}



