<?php
$arr = [];
for ($i = 1; $i <= 5; $i++) {
    $arr[] = intval(trim(fgets(STDIN)));
}
echo max($arr) . PHP_EOL;
echo min($arr) . PHP_EOL;
