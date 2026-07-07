export const keys = {};
window.addEventListener('keydown', function(event) {
    keys[event.key.toLowerCase()] = true;
    console.log("key " + event.key);
});
window.addEventListener('keyup', function(event) {
    keys[event.key.toLowerCase()] = false;
});