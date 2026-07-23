const countdownTimer = document.getElementById("countdownTimer");

export function changeCountdownTimer(timeLeft) {
    if (!Number.isFinite(timeLeft)) {
        console.log("В функцию изменения таймера передано не число:", timeLeft);
        return;
    }

    countdownTimer.textContent = Math.ceil(timeLeft).toString();
}