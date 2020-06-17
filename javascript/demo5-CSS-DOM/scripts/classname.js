
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

// set the new class for tag
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

function styleHeaderSiblings() {
    if(!document.getElementsByTagName) return false;
    var headers = document.getElementsByTagName("h1");
    headers[0].className = "intro"
}

addLoadEvent(styleHeaderSiblings)

