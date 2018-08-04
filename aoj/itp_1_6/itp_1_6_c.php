<?php


/**
 * n: entry
 *
 * b: buildings (max:4)
 * f: floors (max:3)
 * r: rooms (max:10)
 * v: person
 */


# Initialize
for ($b = 1; $b <= 4; $b++) {
    for ($f = 1; $f <= 3; $f++) {
        $house[$b][$f] = array_fill(0, 10, 0);
    }
}

# Input
$n = trim(fgets(STDIN));
for ($i = 0; $i < $n; $i++) {
    $input = explode(" ", trim(fgets(STDIN)));
    $house[$input[0]][$input[1]][$input[2]] = $input[3];
}

# Output
for ($b = 1; $b <= 4; $b++) {
    for ($f = 1; $f <= 3; $f++) {
        echo(" ".implode(" ", $house[$b][$f])."\n");
    }
    if ($b < 4) {
        echo("####################\n");
    }
}
