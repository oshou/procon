<?php

list($r, $c) = explode(" ", trim(fgets(STDIN)), 2);

for ($i = 0; $i < $r; $i++) {
    $row = explode(" ", trim(fgets(STDIN)));

    $sum = 0;
    foreach ($row as $j => $n) {
        $sum += $n;
        $arr[$j] += $n;
    }
    $row = $sum;
    printf("%s\n", implode(' ', $row));
}

$sum = array_sum($arr);
$arr[] = $sum;
printf("s\n", implode(' ', $arr));
