<?php
$arr = [];
for ($i = 0; $i < 3; $i++) {
    $n = trim(fgets(STDIN));
    if (empty($arr[$n])) {
        $arr[$n] = 1;
    } else {
        $arr[$n]++;
    }
}
foreach ($arr as $k => $v) {
    if ($v >= 2) {
        echo $k . PHP_EOL;
    }
}
