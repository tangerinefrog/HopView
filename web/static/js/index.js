const input = document.getElementById('urlInput')
input.addEventListener('keydown', async (e) => {
    if(e.code === 'Enter') {
        const url = e.target.value;
        if (url != null)  {
            const nodes = await getUrlNodes(url);
            const coords = parseCoordinates(nodes);
            console.log(coords);
            drawRoute(coords);
        }
    }
});

async function getUrlNodes(url){
    try{
        const resp = await fetch(`/api/traceroute?url=${url}`);
        const body = await resp.json();

        if (!resp.ok) {
            throw new Error(`Traceroute endpoint status code: ${resp.status}, body:\n${body}`);
        }

        if(body.endpoints?.length > 0){
            return body.endpoints
        } else {
            return []
        }
    } catch (error) {
        console.error(error.message);
    }
}

function parseCoordinates(nodes){
    return nodes.map(n => {
        return [n.Longitude, n.Latitude]
    });
}