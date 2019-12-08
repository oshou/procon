<?php

fscanf(STDIN, "%d %d", $m, $n);

for ($i = 0; $i < 10; $i++) {
    if ($i >= 1) {
        echo " ";
    }
    echo $m + $n * $i;
}
echo PHP_EOL;
