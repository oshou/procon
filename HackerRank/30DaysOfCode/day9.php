<?php

function factorial(int $n): int
{
    if ($n <= 1) {
        return 1;
    }
    return $n * factorial($n - 1);
}

$n = intval(fgets(STDIN));
print (factorial($n)) . PHP_EOL;
