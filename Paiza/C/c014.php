<?php

fscanf(STDIN, "%d %d", $n, $r);
for ($i = 0; $i < n; $i++) {
    fscanf(STDIN, "%d %d %d", $h, $w, $d);
    $diameter = 2 * $r;
    if ($diameter <= $h && $diameter <= $w && $diameter <= $d) {
        echo $i . "\n";
    }
}
