import {GAME_CONSTANTS, WEAPON_SPRITES} from "../../model/game/storage/gameConstants.js";

const killFeed = document.getElementById("killFeed");

export function updateKillFeed(killsInfo, playersNicknames, player, enemies) {
    killsInfo.forEach((killInfo) => {
        const isPlayer = killInfo.kId === player.id || killInfo.vId === player.id;
        console.log(isPlayer);
        const weaponSpritePath = getWeaponSpritePath(killInfo.kId, player, enemies);

        addToKillFeed(playersNicknames[killInfo.kId], playersNicknames[killInfo.vId], weaponSpritePath, isPlayer);
    })
}

function getWeaponSpritePath(killerId, player, enemies) {
    if (killerId === player.id) {
        return returnWeaponPath(player.pc);
    }

    for (const [id, enemy] of Object.entries(enemies)) {
        if (killerId === id) {
            return returnWeaponPath(enemy.pc);
        }
    }

    return "";
}

function returnWeaponPath(playerType) {
    switch (playerType) {
        case GAME_CONSTANTS.PLAYER_TYPES.PISTOL:
            return WEAPON_SPRITES.PISTOL;
        case GAME_CONSTANTS.PLAYER_TYPES.SHOTGUN:
            return WEAPON_SPRITES.SHOTGUN;
        case GAME_CONSTANTS.PLAYER_TYPES.ROCKET_LAUNCHER:
            return WEAPON_SPRITES.ROCKET_LAUNCHER;
        default:
            return "";
    }
}

function addToKillFeed(killerNickname, victimNickname, weapon, isPlayer) {
    const newKillFeedElement = document.createElement("div");
    if (isPlayer) {
        newKillFeedElement.className = "kill-feed-element me";
    } else {
        newKillFeedElement.className = "kill-feed-element";
    }

    const killerInfo = document.createElement("p");
    killerInfo.className = "kill-feed-player-text";
    killerInfo.innerHTML = killerNickname;
    const gunImage = document.createElement("img");
    gunImage.className = "kill-feed-gun-image";
    gunImage.src = weapon;
    gunImage.alt = "";
    const victimInfo = document.createElement("p");
    victimInfo.className = "kill-feed-player-text";
    victimInfo.innerHTML = victimNickname;

    newKillFeedElement.appendChild(killerInfo);
    newKillFeedElement.appendChild(gunImage);
    newKillFeedElement.appendChild(victimInfo);

    killFeed.appendChild(newKillFeedElement);

    setTimeout(() => {
        newKillFeedElement.classList.add("show");
    }, 20);

    const deleteTime = GAME_CONSTANTS.KILL_FEED_DELETE_TIME;
    const transitionTime = 400;

    setTimeout(() => {
        newKillFeedElement.classList.remove("show");
    }, deleteTime - transitionTime);

    setTimeout(() => {
        newKillFeedElement.remove();
    }, deleteTime);
}