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