// configuration

const baseURL = "http://localhost:8000/"

// main

function index_addThreadOnMenu( threadInfo) {
    // prepare dom
    let template = document.querySelector("#threadItem");
    let clone = template.content.cloneNode(true);
    // insert data into template
    clone.querySelector(".thread").innerText = threadInfo.title;
    clone.querySelector(".thread").href = baseURL + "thread/" + threadInfo.thread_id + ".html";
    // insert datas finished here! ----------------
    let menu = document.getElementById("threadList");
    menu.appendChild(clone);
}
function showIndex (){
    const request = new Request("top.json");
    fetch(request)
    .then ( (response) => (response.json()))
    .then ( (data) => {
        data.forEach(element => index_addThreadOnMenu(element))
    })
}

window.addEventListener("load", (event) => {showIndex()})
