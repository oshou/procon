<?php
$n = intval(trim(fgets(STDIN)));
$mod = 21 % $n;
if ($mod === 0) {
    echo $n . PHP_EOL;
} else {
    echo $mod . PHP_EOL;
}
