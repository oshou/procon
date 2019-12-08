<?php
$arr = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J"];
fscanf(STDIN, "%d %d %d", $a, $b, $c);
printf("%s\n", join("", array_slice($arr, 0, $a)));
printf("%s\n", join("", array_slice($arr, $a, $b)));
printf("%s\n", join("", array_slice($arr, $a + $b, $c)));
