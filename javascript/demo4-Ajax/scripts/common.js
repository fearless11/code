

function addLoadEvent(func) {
    var oldonload = window.onload;
    if (typeof window.onload != 'function') {
        window.onload=func;
    }else{
        window.onload= function () {
            oldonload();
            func()
        }
    }
}

function getHTTPObject() {
    return new XMLHttpRequest();
}


function getNewContent() {
    var request = getHTTPObject();
    if (request) {
        request.open("GET","http://localhost:8080/cors",true);
        request.onreadystatechange = function () {
           // readyState 0:初始化 1：正在加载 2：加载完毕 3：正在交互 4：完成
            if(request.readyState == 4) {
                alert("Response Received")
                var para = document.createElement("p");
                var txt = document.createTextNode(request.responseText);
                para.appendChild(txt);
                document.getElementById('new').appendChild(para);
            }
        };
        request.send(null)
    }else{
        alert("sorry, your brower does not support XMLHTTPRquest")
    }
    // 因为是异步请求,所以会先弹出 Function Done,之后才是 Response Received （请求需要时间）
    alert("Function Done")
}

addLoadEvent(getNewContent);


