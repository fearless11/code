
function showPic(whichpic) {

    if (!document.getElementById("placeholder")) {
        return false
    }
    // 改元素节点的属性
    var source = whichpic.getAttribute("href");
    var placeholder = document.getElementById("placeholder");
    placeholder.setAttribute("src",source);

    // 改文本节点的内容
    if (document.getElementById("description")) {
        var text = whichpic.getAttribute("title");
        var description = document.getElementById("description");
        description.firstChild.nodeValue = text;
    }
    return true;
}

function prepareGallery() {
    
    // 确保兼容
    if (!document.getElementsByTagName || !document.getElementById) {
        return false
    }
    if (!document.getElementById("imagegallery")) {
        return false
    }

    var gallery = document.getElementById("imagegallery")
    var links = gallery.getElementsByTagName("a")
    for (let i = 0; i < links.length; i++) {
        // 返回false取消onclick的默认行为
        links[i].onclick = function() {
            return showPic(this) ? false : true;
        }
    }
}

// 网页加载完毕会触发一个onload事件，与window对象相关联
// 绑定一个函数 window.onload = prepareGallery
// 绑定多个函数 window.onload = function(){ function1(); function2(); }
// 通用绑定任意个函数
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


addLoadEvent(prepareGallery)