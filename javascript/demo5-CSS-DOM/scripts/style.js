
function stripeTables() {
    if(!document.getElementsByTagName) return false;
    var tables = document.getElementsByTagName('table');
    var odd,rows;

    for (let i = 0; i < tables.length; i++) {
        odd = false;
        rows = tables[i].getElementsByTagName("tr");
        for (let j = 0; j < rows.length; j++) {
            if (odd == true) {
                // 通过style设置背景色
                rows[j].style.backgroundColor = '#ffc';
                odd=false;
            }else{
                odd=true;
            }
        }
    }
}

function highlightRows() {
    if(!document.getElementsByTagName) return false;
    var rows = document.getElementsByTagName("tr");
    for (let i = 0; i < rows.length; i++) {
        rows[i].onmouseover = function () {
            this.style.fontWeight = "bold";
        }
        rows[i].onmouseout = function () {
            this.style.fontWeight = "normal"
        } 
    }
}

window.onload = function() {
    stripeTables()
    highlightRows()
}