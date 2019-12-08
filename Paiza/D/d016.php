<?php

$s = trim(fgets(STDIN));
$n = intval(fgets(STDIN));
echo substr($s, 0, $n) . "\n";
