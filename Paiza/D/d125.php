<?php
$n = intval(fgets(STDIN));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[] = intval(fgets(STDIN));
}
echo array_sum($arr) / count($arr) . "\n";
