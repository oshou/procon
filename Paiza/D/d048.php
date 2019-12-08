<?php

$arr = [];
for ($i = 0; $i < 5; $i++) {
    $arr[] = intval(fgets(STDIN));
    if ($i >= 1) {
        echo ($arr[$i] - $arr[$i - 1]) . "\n";
    }
}
