const input = document.getElementById('targetInput');
input.addEventListener('keydown', async (e) => {
    if (e.code === 'Enter') {
        const target = e.target.value;
        if (target != null) {
            getTargetNodes(target);
        }
    }
});

let previousCoordinates = [];
const nodesQueue = [];

function getTargetNodes(target) {
    reset();

    const es = new EventSource(`/api/traceroute?target=${target}`);

    es.onmessage = function (e) {
        const node = JSON.parse(e.data);
        if (node) {
            drawNode(node);
        }
    };

    es.onerror = function () {
        es.close();
    };
}

function drawNode(node) {
    const coordinates = parseCoordinates(node);
    if (coordinates.length === 0) {
        return;
    }

    if (previousCoordinates.length === 0) {
        
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