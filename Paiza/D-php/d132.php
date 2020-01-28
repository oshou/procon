<?php
$n = intval(trim(fgets(STDIN)));
$s = trim(fgets(STDIN));
echo str_repeat($s, $n) . PHP_EOL;
