
function addLoadEvent(func) {
    var oldonload = window.onload;
    if (typeof window.onload != 'function') {
        window.onload=func;
    }else{
        window.onload= function () {
            oldonload();
            func();
        }
    }
}


function insertAfter(newElement, targetElement) {
    var parent = targetElement.parentNode;
    if ( parent.lastChild == targetElement ) {
        parent.appendChild(newElement);
    }else { 
        parent.insertBefore(newElement,targetElement.neytSibling);
    }
}

function addClass(element,value) {
    if(!element.className){
        element.className = value;
    }else{
        newClassName = element.className;
        newClassName += " ";
        newClassName += value;
        element.className = newClassName;
    }
}

// 当前导航页加强显示  结合Class属性+CSS实现
function highlightPage() {
    if (!document.getElementsByTagName) {
        return false;
    }
    if(!document.getElementById) {
        return false;
    }

    var headers = document.getElementsByTagName('header');
    if (headers.length == 0 ) return false;
    var navs = headers[0].getElementsByTagName('nav');
    if (navs.length == 0) return false;

    var links = navs[0].getElementsByTagName('a')
    var linkurl;
    for (var i = 0; i < links.length; i++) {
        linkurl = links[i].getAttribute("href");
        if (window.location.href.indexOf(linkurl) != -1) {
            links[i].className = "here";
            // 根据为body设置id后在根据不同css设置页面背景
            var linktext = links[i].lastChild.nodeValue.toLocaleLowerCase();
            document.body.setAttribute("id",linktext);
        }
        
    }
}


// 内部导航， 只显示选中的部分，隐藏其他部分
function showSection(id) {
    var sections = document.getElementsByTagName("section");
    for (var i = 0; i < sections.length; i++) {
        if (sections[i].getAttribute("id") != id) {
            sections[i].style.display = "none";
        }else {
            sections[i].style.display = "block";
        }
    }
}
function prepareInternalnav() {
    if(!document.getElementsByTagName) return false;
    if(!document.getElementById) return false;
    var articles = document.getElementsByTagName("article"); 
    // 调试
    // window.alert(articles.length)     
    if (articles.length == 0) return false;
    var navs = articles[0].getElementsByTagName("nav");
    if (navs.length == 0) return false;
    var nav = navs[0];
    var links = nav.getElementsByTagName("a");
    for (var i = 0; i < links.length; i++) {
        var sectionId = links[i].getAttribute("href").split("#")[1];
        if (!document.getElementById(sectionId)) continue;
        links[i].destination = sectionId;
        links[i].onclick = function () {
            showSection(this.destination);
            return false;
        }
    }
}

// 图片 photos
function showPic(whichpic) {
    if (!document.getElementById("placeholder")) {
        return true;
    }
    var source = whichpic.getAttribute("href");
    var placeholder = document.getElementById("placeholder");
    placeholder.setAttribute("src",source);
    if (!document.getElementById("description")) {
        return false;
    }
    if(whichpic.getAttribute("title")) {
        var text = whichpic.getAttribute("title");
    }else {
        var text = "";
    }

    var description = document.getElementById("description");
    if (description.firstChild.nodeType == 3) {
        description.firstChild.nodeValue = text;
    }
    return false;
}
// 创建图片占位符
function preparePlaceholder() {
    if(!document.createElement) return false;
    if(!document.createTextNode) return false;
    if(!document.getElementById) return false;
    if(!document.getElementById("imagegallery")) return false;

    var placeholder = document.createElement("img");
    placeholder.setAttribute("id","placeholder");
    placeholder.setAttribute("src","images/logo.jpg");
    placeholder.setAttribute("alt","my image gallery");

    var description = document.createElement("p");
    description.setAttribute("id","description");
    var desctext = document.createTextNode("choose an image");
    description.appendChild(desctext);

    var gallery = document.getElementById("imagegallery");
    insertAfter(description,gallery);
    insertAfter(placeholder,description);
}
// 设置点击图片时响应事件
function prepareGallery() {
    if (!document.getElementsByTagName) return false;
    if(!document.getElementById)  return false;
    if(!document.getElementById("imagegallery")) return false;

    var gallery = document.getElementById("imagegallery");
    var links = gallery.getElementsByTagName("a");
    for (var i = 0; i < links.length; i++) {
        links[i].onclick = function () {
            return showPic(this);
        }
    }
}


// 表格 table 
// 表格单双行不同颜色 结合CSS的class实现
function stripeTables() {
    if (!document.getElementsByTagName) return false;
    var tables = document.getElementsByTagName("table");
    for (var i = 0; i < tables.length; i++) {
        var odd = false;
        var rows = tables[i].getElementsByTagName("tr");
        for (var j = 0; j < rows.length; j++) {
            if(odd == true){
                addClass(rows[j],"odd");
                odd=false;
            }else{
                odd=true;
            }
        }
    }
} 
// 表格选中行高亮
function highlightRows() {
    if (!document.getElementsByTagName) {
        return false;
    }
    var rows = document.getElementsByTagName("tr");
    for (var i = 0; i < rows.length; i++) {
         rows[i].oldClassName = rows[i].className;
         rows[i].onmouseover = function () {
                addClass(this,"highlight");
         }
         rows[i].onmouseout = function () {
             this.className = this.oldClassName;
         }
    }
}


// form 字段默认取placeholder
function resetFields(whichform) {
    // Modernizr是HTML5的API
    // https://modernizr.com/download?setclasses
    if(!Modernizr.input.placeholder) return;  
    for (var i = 0; i < whichform.elements.length; i++) {
        var element = whichform.elements[i]
        if (element.type == "submit") continue;

        var check = element.placeholder || element.getAttribute("placeholder");
        if (!check) continue;
        // 用户通过tab键或点击表单字段时触发
        // 清空value值, 当前输入值作为value值
        element.onfocus = function () {
            var text = this.placeholder || this.getAttribute("placeholder");
            if (this.value == text) {
                this.className = "";
                this.value = "";
            }
        }
        // 用户把焦点移出表单字段时触发
        // 取placeholder的值当做value值
        element.onblur = function () {
                if(this.value == ""){
                    this.className = "placeholder";
                    this.value = this.placeholder || this.getAttribute("placeholder");
                }
        }
        element.onblur();
    }
}

function isFilled(field) {
    if (field.value.replace(' ','').length == 0) return false;
    var placeholder = field.placeholder || field.getAttribute('placeholder');
    return (field.value != placeholder);
}

function isEmail(field) {
    return (field.value.indexOf("@") != -1 && field.value.indexOf(".") != -1);
}

// 验证表单
function validateForm(whichform) {
    for (var i = 0; i < whichform.elements.length; i++) {
         var element = whichform.elements[i];
        if (element.required == 'required') {
             if (!isFilled(element)) {
                alert("Please fill in the "+ element.name +"field.");
                return false;
            }                                                                                                                                                   

        }
        if (element.type == 'email') {
            if (!isEmail(element)) {
                alert("the is not a valid email address");
                return false;
            }
        }
    }   
    return true;
}

// 提交表单后,ajax拦截返回submit.html页面
function getHTTPObject() {
    return new XMLHttpRequest();
}

// 创建加载元素
function displayAjaxLoading(element) {
    while (element.hasChildNodes()){
        element.removeChild(element.lastChild);
    }
    var content = document.createElement("img");
    content.setAttribute("src","images/loading.gif");
    content.setAttribute("alt","Loading...");
    element.appendChild(content);
}

function submitFormWithAjax(whichform, thetarget) {
    var request = getHTTPObject();
    if (!request) return false;
    // 加载等待响应, 先移除article的所有子元素,替换为加载图片等待
    displayAjaxLoading(thetarget);

    var dataParts = [];
    var element;
    for (var i = 0; i < whichform.elements.length; i++) {
        element = whichform.elements[i];
        // 对URL编码
        dataParts[i] = element.name + '=' + encodeURIComponent(element.value);
    }
    var data = dataParts.join('&');
    request.open('POST', whichform.getAttribute('action'),true);
    request.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    // 异步响应
    request.onreadystatechange = function () {
        if (request.readyState == 4) {
            if (request.status == 200 || request.status == 0) {               
                var submitHTML = "<article><h1>Thanks!</h1><p>Thanks for contacting us. We'll get back to you as soon as we can.</p>";
                // 匹配服务器响应的article元素的文本内容作为当前页article元素内容
                // var matches = request.responseText.match(/<article>([\s\S]+)<\/article>/);
                var matches = submitHTML.match(/<article>([\s\S]+)<\/article>/);
                if ( matches != null && matches.length > 0 ) {
                    thetarget.innerHTML = matches[1];
                }else {
                    thetarget.innerHTML = '<p>Oops, there was an error. sorry.</p>';
                }
            }else {
                thetarget.innerHTML = '<p>'+ request.statusText+'</p>';
            }
        }
    }

    request.send(data);
    return true;
}

function prepareForms() {
    for (var i = 0; i < document.forms.length; i++) {
        var thisform  = document.forms[i];
        resetFields(thisform);
        thisform.onsubmit = function (){
            if(!validateForm(this)) return false;
            var article = document.getElementsByTagName('article')[0];
            // submit提交后被ajax拦截, ajax发送请求后将返回内容替换到当前页面的article元素内
            if(submitFormWithAjax(this,article)) return false;
            return true;
        }
    }
}          



addLoadEvent(highlightPage);
addLoadEvent(prepareInternalnav);

// photos
addLoadEvent(preparePlaceholder);
addLoadEvent(prepareGallery);

// table
addLoadEvent(stripeTables);
addLoadEvent(highlightRows);

// form 
addLoadEvent(prepareForms);