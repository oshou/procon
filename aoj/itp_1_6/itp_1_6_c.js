line = require('fs').readFileSync('/dev/stdin', 'ascii').split('\n');
line = line.slice(1, parseInt(line[0]) + 1);

bld = [];
for (i = 0; i < 4; i++){
    floor = [];
    for (j = 0; j < 3; j++){
        room = [];
        for (k = 0; k < 10; k++){
            room[k] = 0;
        }
        floor.push(room);
    }
    bld.push(floor);
}

line.forEach(function(line){
    arr = line.split(' ').map(Number).map(function(x) { return x-1;});
    bld[arr[0]][arr[1]][arr[2]] += arr[3] + 1;
});

for (i in build){
    floor = bld[i];
    for (j in floor){
        room = floor[j];
        console.log(''. room.join(' '));
    }
    if (i != '3'){
        console.log('####################');
    }
}
