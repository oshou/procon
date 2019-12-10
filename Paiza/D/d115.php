<?php
$n = intval(fgets(STDIN));
if ($n % 2 === 0) {
    printf("%d\n", $n / 2);
} else {
    printf("%d\n", floor($n / 2));
}
