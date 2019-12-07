<?php
$handle = fopen("php://stdin", "r");
$i = 4;
$d = 4.0;
$s = "HackerRank ";

$i2 = fgets($handle);
$d2 = fgets($handle);
$s2 = fgets($handle);

printf("%d\n", $i + $i2);
printf("%.1f\n", $d + $d2);
print($s . $s2 . "\n");

fclose($handle);
