<?php
$x = intval(trim(fgets(STDIN)));
$a = intval(trim(fgets(STDIN)));
$b = intval(trim(fgets(STDIN)));
echo intval($x / $a) * ($b - $a) . PHP_EOL;
