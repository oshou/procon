<?php
$n = intval(trim(fgets(STDIN)));
fscanf(STDIN, "%d %d", $h, $w);
echo (($h * $w) % $n) . PHP_EOL;
