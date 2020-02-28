<?php
fscanf(STDIN, "%d %d", $m, $n);
for ($i = 1; $i <= 10; $i++) {
    if ($i != 1) {
        echo " ";
    }
    echo $m + $n * ($i - 1);
}
echo PHP_EOL;
