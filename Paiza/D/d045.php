<?php

$n = intval(fgets(STDIN));
$map = [
    1 => "E",
    2 => "D",
    3 => "C",
    4 => "B",
    5 => "A",
];
echo $map[$n] . "\n";
