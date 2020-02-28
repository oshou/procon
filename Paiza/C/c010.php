<?php
fscanf(STDIN, "%d %d %d", $a, $b, $r);
$n = intval(fgets(STDIN));
$x = [];
for ($i = 0; $i < $n; $i++) {
    $x[$i] = explode(" ", trim(fgets(STDIN)));
    $d = ((int) $x[$i][0] - $a) ** 2 + ((int) $x[$i][1] - $b) ** 2;
    $area = $r ** 2;
    echo ($d > $area) ? "silent" : "noisy";
    echo PHP_EOL;
}
