<?php

$n = intval(fgets(STDIN));
if ($n < 0) {
    echo (-1) * $n;
} else {
    echo $n;
}
echo PHP_EOL;
