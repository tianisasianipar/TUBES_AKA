let chart;
let labels = [];
let iterData = [];
let recData = [];

window.onload = () => {
    const ctx = document.getElementById('timeChart').getContext('2d');
    chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels,
            datasets: [
                {
                    label: 'Iteratif',
                    data: iterData,
                    borderColor: '#514659',
                    tension: 0.3
                },
                {
                    label: 'Rekursif',
                    data: recData,
                    borderColor: '#ae8c93',
                    tension: 0.3
                }
            ]
        }
    });
};

async function runAnalysis() {
    const n = document.getElementById('inputN').value;
    const res = await fetch(`/api/analyze?n=${n}`);

    if (!res.ok) {
        alert("Input tidak valid!");
        return;
    }

    const data = await res.json();

    document.getElementById("iterTime").innerText = data.iterative.toFixed(4);
    document.getElementById("recTime").innerText = data.recursive.toFixed(4);

    labels.push(data.n);
    iterData.push(data.iterative);
    recData.push(data.recursive);

    chart.update();

    addTableRow(data.n, data.iterative, data.recursive);
}

//
function addTableRow(n, iter, rec) {
    const tableBody = document.querySelector("#resultTable tbody");

    const row = document.createElement("tr");

    row.innerHTML = `
        <td>${n}</td>
        <td>${iter.toFixed(4)}</td>
        <td>${rec.toFixed(4)}</td>
    `;

    tableBody.appendChild(row);
}

