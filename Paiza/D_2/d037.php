<?php
$m = intval(trim(fgets(STDIN)));
$n = intval(trim(fgets(STDIN)));
if ($n === 0) {
    exit;
} else {
    echo intval(ceil($n / $m)) . PHP_EOL;
}
