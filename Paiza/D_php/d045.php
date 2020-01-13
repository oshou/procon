<?php
$map = [
    1 => "E",
    2 => "D",
    3 => "C",
    4 => "B",
    5 => "A",
];
$n = intval(trim(fgets(STDIN)));
echo $map[$n] . PHP_EOL;
