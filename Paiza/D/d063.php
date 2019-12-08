<?php

$arr = explode(" ", trim(fgets(STDIN)));
sort($arr);

$n = trim(fgets(STDIN));
for ($i = 0; $i < count($arr); $i++) {
    if ($arr[$i] <= $n) {
        continue;
    }
    printf("%d\n", $i + 1);
    exit;
}
printf("%d\n", 6);
