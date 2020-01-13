<?php
$str = "ABCDEFGHIJ";
fscanf(STDIN, "%d %d %d", $s1, $s2, $s3);
echo substr($str, 0, $s1) . PHP_EOL;
echo substr($str, $s1, $s2) . PHP_EOL;
echo substr($str, $s1 + $s2, $s3) . PHP_EOL;
