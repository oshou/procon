<?php
$n = (int) fgets(STDIN);
for ($i = 1; $i <= 9; $i++) {
    if ($i != 1) {
        echo " ";
    }
    echo $n * $i;
}
echo PHP_EOL;
