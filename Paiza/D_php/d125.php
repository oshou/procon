<?php
$n = intval(fgets(STDIN));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[] = intval(trim(fgets(STDIN)));
}
echo floor(array_sum($arr) / count($arr)) . PHP_EOL;
