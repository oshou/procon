<?php
fscanf(STDIN, "%d %d", $n, $m);
if ($n + $m < 16) {
    echo "HIT" . PHP_EOL;
} else {
    echo "STAND" . PHP_EOL;
}
