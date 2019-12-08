<?php
fscanf(STDIN, "%d %d", $a, $b);
$m = $a * $b;
if ($m >= 10000) {
    print("NG\n");
} else {
    printf("%d\n", $m);
}
