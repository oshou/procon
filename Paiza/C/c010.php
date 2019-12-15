<?php

fscanf(STDIN, "%d %d %d", $a, $b, $R);
$n = intval(fgets(STDIN));
for ($i = 0; $i < $n; $i++) {
    fscanf(STDIN, "%d %d", $x, $y);
    $judge = ($x - $a) ** 2 + ($y - $b) ** 2;
    if ($judge >= $R ** 2) {
        echo "silent\n";
    } else {
        echo "noisy\n";
    }
}
