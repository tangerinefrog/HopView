const map = new maplibregl.Map({
    container: 'map',
    style: 'https://basemaps.cartocdn.com/gl/voyager-gl-style/style.json',
    zoom: 3,
    attributionControl: false
});

let markers = [];
const linePrefix = 'line-';

map.on('style.load', () => {
    map.setProjection({
        type: 'globe',
    });
});

function addMarker(coords, text) {
    const markerEl = document.createElement('div');
    markerEl.className = 'custom-marker';

    const popup = new maplibregl.Popup({
        closeButton: false,
    })
        .setHTML(text);

    const marker = new maplibregl.Marker(markerEl)
        .setLngLat(coords)
        .setPopup(popup)
        .addTo(map);

    markers.push(marker);
}

function addLine(start, end) {
    const id = linePrefix + crypto.randomUUID();

    map.addSource(id, {
        type: 'geojson',
        data: {
            type: 'Feature',
            geometry: {
                type: 'LineString',
                coordinates: [start, end],
            },
        },
    });

    map.addLayer({
        id: id,
        type: 'line',
        source: id,
        layout: {
            'line-join': 'round',
            'line-cap': 'round',
        },
        paint: {
            'line-color': '#ff0000',
            'line-width': 3,
        },
    });
}

function clearMap() {
    if (markers) {
        markers.forEach(m => m.remove());
        markers = [];
    }

    const lineLayers = map.getStyle().layers?.filter(l => l.id.startsWith(linePrefix));
    if (lineLayers) {
        lineLayers.forEach(l => {
            if (map.getLayer(l.id)) {
                map.removeLayer(l.id);
            }
            if (map.getSource(l.id)) {
                map.removeSource(l.id);
            }
        });
    }
}

function flyTo(coordinates) {
    map.flyTo({
        center: coordinates,
        essential: true,
    })
}