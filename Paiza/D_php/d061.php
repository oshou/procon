<?php
$c = intval(trim(fgets(STDIN)));
if ($c === 0) {
    printf("%d\n", 1);
} else {
    printf("%d\n", $c * 3);
}
