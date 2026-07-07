const stateDirection = {
    DOWN: 'down',
    UP: 'up',
    LEFT: 'left',
    RIGHT: 'right',
    DOWN_LEFT: 'down_left',
    DOWN_RIGHT: 'down_right',
    UP_LEFT: 'up_left',
    UP_RIGHT: 'up_right',
    NONE: 'none',
}
let direction = stateDirection.NONE;
export const keys = {};
window.addEventListener('keydown', function(event) {
    keys[event.key.toLowerCase()] = true;
});
window.addEventListener('keyup', function(event) {
    keys[event.key] = false;
});