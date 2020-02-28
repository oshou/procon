<?php
$n = intval(trim(fgets(STDIN)));
$d = [];
$sum = 0;
for ($i = 1; $i <= $n; $i++) {
    $d = explode(" ", trim(fgets(STDIN)));
    $multiple = 0.01;
    if (substr_count($d[0], "5") > 0) {
        $multiple = 0.05;
    } else if (substr_count($d[0], "3") > 0) {
        $multiple = 0.03;
    }
    $point = intval($d[1] * $multiple);
    $sum += $point;
}
echo $sum . PHP_EOL;
