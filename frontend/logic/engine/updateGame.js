export function updateGame(elapsedTime, state) {
    //TODO: написать игру =)

    //Отладочный код для перемещения квадратика
    state.square.x += 0.01 * elapsedTime;
    state.square.y += 0.005 * elapsedTime;
}