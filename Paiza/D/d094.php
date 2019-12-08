<?php

$arr = [];
for ($i = 0; $i < 3; $i++) {
    $n = trim(fgets(STDIN));
    if (isset($arr[$n])) {
        $arr[$n] += 1;
    } else {
        $arr[$n] = 1;
    }
}
printf("%s\n", array_keys($arr, max($arr))[0]);
