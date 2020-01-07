<?php

const MAX_SIZE = 6;
const MIN_SUM = -63;

for ($i = 0; $i < MAX_SIZE; $i++) {
    $arr[$i] = explode(" ", trim(fgets(STDIN)));
}

$max_sum = MIN_SUM;
for ($i = 0; $i < 4; $i++) {
    for ($j = 0; $j < 4; $j++) {
        $sum = $arr[$i][$j] + $arr[$i][$j + 1] + $arr[$i][$j + 2] + $arr[$i + 1][$j + 1] + $arr[$i + 2][$j] + $arr[$i + 2][$j + 1] + $arr[$i + 2][$j + 2];
        if ($sum > $max_sum) {
            $max_sum = $sum;
        }
    }
}
print ($max_sum) . PHP_EOL;
