<?php
$arr = [];
for ($i = 0; $i < 5; $i++) {
    $arr[] = intval(trim(fgets(STDIN)));
}

for ($i = 0; $i < 5; $i++) {
    if ($i > 0) {
        echo ($arr[$i] - $arr[$i - 1]) . PHP_EOL;
    }
}
