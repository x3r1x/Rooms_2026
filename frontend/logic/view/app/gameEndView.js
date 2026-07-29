const resultList = document.getElementById("resultList");

export function fillResultWindow(ownId, nicknamesList, resultsList) {
    resultsList.sort((result1, result2) => statisticComparator(result1, result2))

    let i = 0;
    resultList.innerHTML = resultsList.map(result => {
        i += 1;
        const isPlayer = result.id === ownId;
        const playerStyle = `style="background-color: #dfff00b0"`;
        const KD = result.d !== 0 ? (result.k / result.d).toFixed(2) : "-"

        return `
            <div class="result-list-element" ${isPlayer ? playerStyle : ""}>
                <p>${i}. ${nicknamesList[result.id]}</p>
                <p>Kills: ${result.k}</p>
                <p>Deaths: ${result.d}</p>
                <p>K/D: ${KD}</p>
            </div>`
    }).join('')
}

function statisticComparator(stat1, stat2) {
    if (stat1.k === stat2.k) {
        if (stat1.d === stat2.d) {
            return stat1.id.localeCompare(stat2.id);
        }

        return stat1.d - stat2.d;
    }

    return stat2.k - stat1.k;
}