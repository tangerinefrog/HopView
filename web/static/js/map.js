const map = new maplibregl.Map({
    container: 'map',
    style: 'https://basemaps.cartocdn.com/gl/voyager-gl-style/style.json',
    zoom: 3,
    attributionControl: false
});

map.on('style.load', () => {
    map.setProjection({
        type: 'globe',
    });
});

function drawRoute(coordinates){
    const route = {
        'type': 'FeatureCollection',
        'features': [
            {
                'type': 'Feature',
                'geometry': {
                    'type': 'LineString',
                    'coordinates': coordinates
                }
            }
        ]
    };

    const source = map.getSource('route');
    if (source){
        map.removeLayer('route');
        map.removeSource('route');
    }

    map.addSource('route', {
        'type': 'geojson',
        'data': route
    });

    map.addLayer({
        'id': 'route',
        'source': 'route',
        'type': 'line',
        'paint': {
            'line-width': 2,
            'line-color': '#007cbf'
        }
    });
}