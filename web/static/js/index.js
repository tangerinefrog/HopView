const input = document.getElementById('urlInput')
input.addEventListener('keydown', (e) => {
    if(e.code === 'Enter') {
        const url = e.target.value;
        if (url != null)  {
            const nodes = getUrlNodes(url)
        }
    }
});

function getUrlNodes(url){
    fetch(`/api/traceroute?url=${url}`)
        .then(resp => resp.json())
        .then(json => console.log(json));
}