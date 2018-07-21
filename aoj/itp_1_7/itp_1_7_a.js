var input = require('fs').readFileSync('/dev/stdin', 'utf8');
var Array = (input.trim()).split(" ");

for (var i=0; i < Array.length; i++){
    var m = Array[0];
    var f = Array[1];
    var r = Array[2];
    judge(m, f, r);
}

function judge(m, f, r){
    sum = m + f;
    if (m == -1 && f == -1 && r == -1) grade = 'F';
    else if (sum >= 80) grade = 'A';
    else if (sum >= 65) grade = 'B';
    else if (sum >= 50) grade = 'C';
    else if (sum >= 30){
        if (r >= 50) grade = 'C';
        else grade = 'D';
    }
    else grade = 'F';

    console.log(grade);
}
