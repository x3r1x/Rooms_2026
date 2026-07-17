import {GAME_CONSTANTS, GAME_SPRITES} from "../../model/storage/gameConstants.js";
import {keys} from '../../controller/listeners.js'
export function drawPlayers(context, mainPlayer, enemies) {
    drawMainPlayer(context, mainPlayer);

    enemies.forEach((enemy) => drawEnemy(context, enemy));
}

function drawMainPlayer(context, player) {
    const spriteSheet = GAME_SPRITES.PLAYER_GOES;
    const frameWidth = spriteSheet.width / 6;
    const frameHeight = spriteSheet.height;
    let frameIndex = 0;
    const size = GAME_CONSTANTS.PLAYER_VISUAL_SIZE;
    if (keys['w'] || keys['a'] || keys['s'] || keys['d'] ||
        keys['ц'] || keys['ф'] || keys['ы'] || keys['в']){
        frameIndex = Math.floor(Date.now() / 100) % 6;
    } else {
        frameIndex = 0;
    }
    context.save();
    context.translate(player.x, player.y);
    context.rotate(player.direction);
    context.drawImage(
        spriteSheet,
        frameIndex * frameWidth, 0, frameWidth, frameHeight,
        -frameWidth / 2, -frameHeight / 2, frameWidth, frameHeight
    );
    context.restore();
}

function drawEnemy(context, enemy) {
    const sprite = GAME_SPRITES.ENEMY_GOES;
    const size = GAME_CONSTANTS.PLAYER_VISUAL_SIZE;

    context.save();
    context.translate(enemy.x, enemy.y);
    context.rotate(enemy.direction);
    context.drawImage(sprite, -size / 2, -size / 2, size, size);
    context.restore();
}