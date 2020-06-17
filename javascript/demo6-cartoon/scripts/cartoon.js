
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

function moveElement(elementID, final_x,final_y,interval) {
    if(!document.getElementById) return false;
    if(!document.getElementById(elementID)) return false;

    var elem = document.getElementById(elementID);
    if(elem.movement) {
        clearTimeout(elem.movement)
    } 

    if (!elem.style.left) {
        elem.style.left = "0px";
    }
    if (!elem.style.top) {
        elem.style.top = "0px";
    }

    var xpos = parseInt(elem.style.left);
    var ypos = parseInt(elem.style.top);
    var dist = 0;
    if(xpos == final_x && ypos == final_y) {
        return true;
    }

    if (xpos < final_x) {
        dist = Math.ceil((final_x - xpos)/10);
        xpos = xpos + dist;
    }
    if (xpos > final_x) {
        dist = Math.ceil((final_x - xpos)/10);
        xpos = xpos - dist;
    }
    if (ypos < final_y) {
        dist = Math.ceil((final_y - ypos)/10);
        ypos = ypos + dist;
    }
    if (ypos > final_y) {
        dist = Math.ceil((final_y - ypos)/10);
        ypos = ypos - dist;
    }

    elem.style.left = xpos + "px";
    elem.style.top = ypos + "px";
    var repeat = "moveElement('" + elementID + "'," + final_x + "," + final_y + "," + interval +")"
    elem.movement = setTimeout(repeat,interval);
}


function insertAfter(newElement, targetElement) {
    var parent = targetElement.parentNode;
    if ( parent.lastChild == targetElement ) {
        parent.appendChild(newElement);
    }else { 
        parent.insertBefore(newElement,targetElement.neytSibling);
    }
}

function prepareSlideshow() {
    if(!document.getElementsByTagName) return false;
    if(!document.getElementById) return false;
    if(!document.getElementById("linklist")) return false;

    var slideshow = document.createElement("div");
    slideshow.setAttribute("id","slideshow");
    var preview = document.createElement("img");
    preview.setAttribute("src","images/1.jpg");
    preview.setAttribute("alt","web cartoon");
    preview.setAttribute("id","preview");
    slideshow.appendChild(preview);
    var list = document.getElementById("linklist");
    insertAfter(slideshow,list);

    // CSS设置图片显示大小,超过部分溢出不显示,每次移动达到轮播的效果
    var links = list.getElementsByTagName("a");
    links[0].onmouseover = function() {
        moveElement("preview",-100,0,10);     
    }
    links[1].onmouseover = function() {
        moveElement("preview",-200,0,10);     
    }
    links[2].onmouseover = function() {
        moveElement("preview",-300,0,10);     
    }
}

addLoadEvent(prepareSlideshow)
addLoadEvent(moveElement)