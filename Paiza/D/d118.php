<?php
$s = trim(fgets(STDIN));
$n = intval(fgets(STDIN));
if ($s == "S") {
    printf("%d\n", (1925 + $n));
} else {
    printf("%d\n", (1988 + $n));
}
