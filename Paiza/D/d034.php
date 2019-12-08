<?php

$n = intval(fgets(STDIN));
$ans = 21 % $n;
if ($ans === 0) {
    echo $n . "\n";
} else {
    echo $ans;
}
