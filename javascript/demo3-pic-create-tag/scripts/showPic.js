
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


// 创建节点标记插入到已有的HTML文档节点之后
function insertAfter(newElement, targetElement) {
    var parent = targetElement.parentNode;
    if ( parent.lastChild == targetElement ) {
        parent.appendChild(newElement);
    }else { 
        parent.insertBefore(newElement,targetElement.nextSibling);
    }
}

// 创建节点tag
function preparePlaceholder() {
    if(!document.createElement) return false;
    if(!document.createTextNode) return false;
    if(!document.getElementById) return false;
    if(!document.getElementById("imagegallery")) return false;

    var placeholder = document.createElement("img");
    placeholder.setAttribute("id","placeholder");
    placeholder.setAttribute("src","images/1.jpg");
    placeholder.setAttribute("alt","my image gallery");

    var description = document.createElement("p");
    description.setAttribute("id","description");
    var desctext = document.createTextNode("choose an image");
    description.appendChild(desctext);
    var gallery = document.getElementById("imagegallery");

    //  效果等价于在ul节点后增加img和p节点
    // <img id="placeholder" src="images/1.jpg" alt="my image gallery">
    // <p id="description" class="warning">Choose an image</p>
    insertAfter(placeholder,gallery)
    insertAfter(desctext,placeholder)
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


addLoadEvent(preparePlaceholder);
addLoadEvent(prepareGallery);