
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


// 鼠标停留在某个表格上方时,把该行文本加黑加粗
function highlightRows() {
    if (!document.getElementsByTagName) {
        return false;
    }

    var rows = document.getElementsByTagName("tr");
    for (var i = 0; i < rows.length; i++) {
        rows[i].onmouseover = function () {
            this.style.fontWeight = "bold";
        }
        rows[i].onmouseout = function() {
            this.style.fontWeight = "normal"
        }
    }
}

addLoadEvent(highlightRows)