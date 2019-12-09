<?php

$n = intval(fgets(STDIN));
if ($n < 10) {
    printf("%d\n", 1000);
} else {
    printf("%d\n", 150 * $n);
}
