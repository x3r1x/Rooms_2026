const resultList = document.getElementById("resultList");

export function fillResultWindow(ownId, resultsList) {
    let i = 0;
    resultList.innerHTML = resultsList.map(result => {
        i += 1;
        const isPlayer = result.id === ownId;
        const playerStyle = `style="background-color: #dfff00b0"`;

        return `<p class="result-list-element" ${isPlayer ? playerStyle : ""}>${i}. ${result.n}${isPlayer ? "(You!)" : ""}: ${result.k} / ${result.d}!</p>`
    }).join('')
}