

/* application configuration */
const config = {
    "viewerURL":"http://localhost:8000",
    "apiURL":"https://golang.mipsparc.net"
}
// ==================================================================
function index_addThreadOnMenu( threadInfo) {
    // prepare dom
    let template = document.querySelector("#threadItem");
    let clone = template.content.cloneNode(true);
    // insert data into template
    clone.querySelector(".thread").innerText = threadInfo.title;

    clone.querySelector(".thread").href = config.viewerURL + "/thread/" + threadInfo.thread_id + ".html";
    // insert datas finished here! ----------------
    let menu = document.getElementById("threadList");
    menu.appendChild(clone);
}
function showIndex (){
    const request = new Request(config.apiURL);
    fetch(request)
    .then ( (response) => (response.json()))
    .then ( (data) => {
        data.forEach(element => index_addThreadOnMenu(element))
    })
}

// ==================================================================




// ==================================================================
/* PATHと処理を紐付けるモジュール*/
const pathBinder = {
    bindingTable:[
        // path : Javascript location pathname をLookupする想定。先頭の有効な文字は / である。
        // func : 関数をバインドする
        {"path": /^\/thread\/\d+$/, "func":"Thread Viewer"},
        {"path": /^\/$/, "func": showIndex}
    ],
    // lookup : path -> function or none
    lookup:function(pathKey) {
        let filterResult = this.bindingTable.filter(
            (bindItem) => {
                return bindItem.path.test(pathKey)
            }
        )
        if (filterResult.length == 0) {
            // @todo fallback処理をここに書く 
            filterResult.push = {"path": /^\/$/, "func": showIndex}

        }
        return filterResult[0].func;
    }
};
// ==================================================================
// VIEWER
function viewer(){
    let path = location.pathname;
    console.log("current path:",path);
    let hoge = pathBinder.lookup(path);
    hoge();
}
// ---
// application startup routine
window.addEventListener("load", (event) => {viewer()})
