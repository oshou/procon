<?php
$n = intval(trim(fgets(STDIN)));
if ($n < 10) {
    echo 1000 . PHP_EOL;
} else {
    echo ($n * 150) . PHP_EOL;
}
