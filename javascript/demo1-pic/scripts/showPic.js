

// DOM 文档对象模型
function showPic(whichpic) {
    // 改元素节点的属性
    var source = whichpic.getAttribute("href");
    var placeholder = document.getElementById("placeholder");
    placeholder.setAttribute("src",source);

    // 改文本节点的内容
    var text = whichpic.getAttribute("title");
    var description = document.getElementById("description");
    description.firstChild.nodeValue = text;
}

// BOM 浏览器对象模型 
function popUp(winURL) {
    window.open(winURL,"popup","width=320,height=480")
}