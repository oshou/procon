<?php
fscanf(STDIN, "%d %d", $a, $b);
$n = intval(trim(fgets(STDIN)));

$ans = "";
for ($i = 1; $i <= $n; $i++) {
    // Input
    fscanf(STDIN, "%d %d", $da, $db);
    // Judge
    if ($a === $da) {
        $ans = ($b < $db) ? "High" : "Low";
    } else {
        $ans = ($a > $da) ? "High" : "Low";
    }
    echo $ans . PHP_EOL;
}
