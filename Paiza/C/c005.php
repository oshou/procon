<?php

$n = intval(fgets(STDIN));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[] = explode(".", trim(fgets(STDIN)));
}

for ($i = 0; $i < $n; $i++) {
    $judge = "True";
    foreach ($arr[$i] as $ele) {
        $e = intval($ele);
        if ($e < 0 || 255 < $e) {
            $judge = "False";
            continue;
        }
    }
    printf("%s\n", $judge);
}
