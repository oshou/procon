<?php
$n = intval(fgets(STDIN));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[] = trim(fgets(STDIN));
}
foreach ($arr as $key => $value) {
    printf("%s", $value);
    if ($key === count($arr) - 1) {
        print("\n");
    } else {
        print(" ");
    }
}
