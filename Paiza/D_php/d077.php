<?php
fscanf(STDIN, "%d %d", $a, $b);
$multiple = $a * $b;

if ($a <= 9999) {
    echo $multiple . PHP_EOL;
} else {
    echo "NG" . PHP_EOL;
}
