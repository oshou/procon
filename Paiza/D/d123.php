<?php
$x = intval(fgets(STDIN));
if ($x < 10000) {
    printf("%d\n", $x + 10000);
} else {
    printf("%d\n", $x);
}
