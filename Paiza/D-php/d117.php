<?php
$n = intval(fgets(STDIN));
$arr = explode(" ", trim(fgets(STDIN)));
echo array_sum($arr) . PHP_EOL;
