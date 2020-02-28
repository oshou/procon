<?php
fscanf(STDIN, "%d %d", $n, $r);
$diameter = 2 * $r;
$h = [];
$w = [];
$d = [];
for ($i = 1; $i <= $n; $i++) {
    fscanf(STDIN, "%d %d %d", $h, $w, $d);
    if ($diameter <= $h && $diameter <= $w && $diameter <= $d) {
        echo $i . PHP_EOL;
    }
}
