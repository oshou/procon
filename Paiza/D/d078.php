<?php

$arr = explode(" ", trim(fgets(STDIN)));
$x = intval(fgets(STDIN));
$avg = array_sum($arr) / count($arr);
if ($avg >= $x) {
    print("pass\n");
} else {
    print("failure\n");
}
