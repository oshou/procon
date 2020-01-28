<?php
$n = intval(trim(fgets(STDIN)));
$input = [];
for ($i = 0; $i < $n; $i++) {
    $input[$i] = explode(" ", trim(fgets(STDIN)));
}
$m = intval(trim(fgets(STDIN)));
$pattern = [];
for ($j = 0; $j < $m; $j++) {
    $pattern[$j] = explode(" ", trim(fgets(STDIN)));
}

for ($i = 0; $i <= $n - $m; $i++) {
    for ($j = 0; $j <= $n - $m; $j++) {
    }
}
