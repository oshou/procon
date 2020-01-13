<?php
$n = intval(trim(fgets(STDIN)));
if ($n % 2 == 0) {
    echo "OFF" . PHP_EOL;
} else {
    echo "ON" . PHP_EOL;
}
