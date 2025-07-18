const input = document.getElementById('targetInput');
const loader = document.getElementById('inputLoader');
const errorBox = document.getElementById('errorMessage');

const nodesQueue = [];
let previousCoordinates = [];

input.addEventListener('keydown', async (e) => {
    if (e.code === 'Enter') {
        hideError();

        const target = e.target.value?.trim();
        if (!isValidUrlOrIP(target)) {
            showError('Please enter a valid URL or IP address.');
            return;
        }

        getTargetNodes(target);
    }
});

function getTargetNodes(target) {
    reset();
    toggleLoader(true);

    const es = new EventSource(`/api/traceroute?target=${target}`);

    es.onmessage = function (e) {
        const node = JSON.parse(e.data);
        if (node) {
            drawNode(node);
        }
    };

    es.addEventListener("done", function () {
        toggleLoader(false);
        es.close();
    });

    es.onerror = function () {
        toggleLoader(false);
        es.close();
    };
}

function drawNode(node) {
    const coordinates = parseCoordinates(node);
    if (coordinates.length === 0) {
        return;
    }

    if (previousCoordinates.length === 0) {
        flyTo(coordinates);
    } else {
        addLine(previousCoordinates, coordinates);
    }

    addMarker(coordinates, node.ip);

    previousCoordinates = coordinates;
}

function reset() {
    previousCoordinates = [];
    clearMap();
}

function parseCoordinates(node) {
    return [node.longitude, node.latitude];
}

function isValidUrlOrIP(value) {
    if(!value){
        return false;
    }

    const urlRegex = /^(https?:\/\/)?([\w\-]+\.)+[\w\-]{2,}(\/\S*)?$/i;
    const ipRegex = /^(\d{1,3}\.){3}\d{1,3}$/;
    return urlRegex.test(value) || ipRegex.test(value);
}

function showError(message) {
    errorBox.textContent = message;
    errorBox.style.display = 'block';
}

function hideError() {
    errorBox.textContent = '';
    errorBox.style.display = 'none';
}

function toggleLoader(show) {
    loader.style.display = show ? 'block' : 'none';
    input.disabled = show;
}