<?php

while (fscanf(STDIN, "%d %d", $n, $x)) {
    if ($n == 0 && $x == 0) {
        break;
    }
    $cnt = 0;
    for ($i = 1; $i <= $n; $i++) {
        for ($j = $i+1; $j <= $n; $j++) {
            $c = $x - $i - $j;
            if ($c > $j && $c <= $n) {
                $cnt++;
            }
        }
    }
}
