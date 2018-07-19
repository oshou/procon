<?php

while (($line = trim(fgets(STDIN))) !== '') {
    sscanf($line, '%d%d%d', $m, $f, $r);
    $sum = $m + $f;

    if ($m < 0 && $f < 0 && $r < 0) {
        break;
    } elseif ($sum < 30 || $m < 0 || $f < 0) {
        echo "F\n";
    } elseif ($sum<50) {
        echo $r >= 50 ? "C\n" : "D\n";
    } elseif ($sum<65) {
        echo "C\n";
    } elseif ($sum<80) {
        echo "B\n";
    } else {
        echo "A\n";
    }
}
