<?php
$n = intval(trim(fgets(STDIN)));
$arr = explode(" ", trim(fgets(STDIN)));
echo $arr[$n - 1] . PHP_EOL;
