<?php
$arr = [1, 2, 3, 4, 5];
for ($i = 0; $i < 4; $i++) {
    $num = intval(fgets(STDIN));
    $index = array_search($num, $arr);
    array_splice($arr, $index, 1);
}
printf("%d\n", $arr[0]);
