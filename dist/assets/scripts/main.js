// Manage Excel template download URL
function downloadTemplate() {
    window.location.href = '/download-template';
}

document.addEventListener('DOMContentLoaded', () => {
    const dropzone = document.getElementById('dropzone');
    const fileInput = document.getElementById('fileInput');
    const chartCanvas = document.getElementById('chart');
    let chart;

    // Manage drag-and-drop events
    dropzone.addEventListener('dragover', (e) => {
        e.preventDefault();
        dropzone.classList.add('dragover');
    });

    dropzone.addEventListener('dragleave', () => {
        dropzone.classList.remove('dragover');
    });

    dropzone.addEventListener('drop', (e) => {
        e.preventDefault();
        dropzone.classList.remove('dragover');
        const files = e.dataTransfer.files;
        if (files.length) {
            handleFileUpload(files[0]);
        }
    });

    dropzone.addEventListener('click', () => {
        fileInput.click();
    });

    fileInput.addEventListener('change', (e) => {
        if (fileInput.files.length) {
            handleFileUpload(fileInput.files[0]);
        }
    });

    // Manage File download
    function handleFileUpload(file) {
        const formData = new FormData();
        formData.append('file', file);

        fetch('/upload', {
            method: 'POST',
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            if (data.dates) {
                renderChart(data);
            } else {
                alert('Erreur lors du téléchargement du fichier');
            }
        })
        .catch(error => {
            console.error('Erreur:', error);
            alert('Erreur lors du téléchargement du fichier');
        });
    }

    function renderChart(data) {
        if (chart) {
            chart.destroy();
        }

        chart = new Chart(chartCanvas, {
            type: 'line',
            data: {
                labels: data.dates,
                datasets: [
                    {
                        label: 'Revenues',
                        data: data.revenues,
                        borderColor: 'green',
                        fill: false
                    },
                    {
                        label: 'Expenses',
                        data: data.expenses,
                        borderColor: 'red',
                        fill: false
                    },
                    {
                        label: 'Net Difference',
                        data: data.netDifference,
                        borderColor: 'blue',
                        fill: false
                    }
                ]
            },
            options: {
                responsive: true,
                scales: {
                    x: {
                        type: 'time',
                        time: {
                            unit: 'day'
                        }
                    },
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
    }
});
