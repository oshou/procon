<?php
$s = trim(fgets(STDIN));
$n = intval(trim(fgets(STDIN)));
echo substr($s, 0, $n) . PHP_EOL;
