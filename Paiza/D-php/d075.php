<?php
$arr = [1, 2, 3, 4, 5];
for ($i = 0; $i < 4; $i++) {
    $n = intval(trim(fgets(STDIN)));
    $idx = array_search($n, $arr);
    unset($arr[$idx]);
}
echo array_values($arr)[0] . PHP_EOL;
