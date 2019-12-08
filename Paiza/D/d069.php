<?php

$arr = explode(" ", trim(fgets(STDIN)));
echo round(array_sum($arr) / count($arr), 1) . "\n";
