<?php

$arr = [];
for ($i = 0; $i < 5; $i++) {
    $arr[] = intval(fgets(STDIN));
}
echo max($arr) . "\n";
echo min($arr) . "\n";
