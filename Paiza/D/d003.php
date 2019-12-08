<?php

$n = intval(fgets(STDIN));
for ($i = 1; $i <= 9; $i++) {
    if ($i >= 2) {
        echo " ";
    }
    echo $n * $i;
}
echo PHP_EOL;
