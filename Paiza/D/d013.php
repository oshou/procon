<?php
fscanf(STDIN, "%d %d", $m, $n);
if ($n !== 0) {
    printf("%d %d", $m / $n, $m % $n);
} else {
    exit;
}
