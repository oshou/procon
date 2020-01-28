<?php
$scores = explode(" ", trim(fgets(STDIN)));
$n = intval(trim(fgets(STDIN)));
$avg = array_sum($scores) / count($scores);
echo ($avg >= $n) ? "pass" : "failure";
echo PHP_EOL;
