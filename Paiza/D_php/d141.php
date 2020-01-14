<?php
$n = intval(trim(fgets(STDIN)));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[$i] = trim(fgets(STDIN));
}
echo implode(" ", $arr) . PHP_EOL;
