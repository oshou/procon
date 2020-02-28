<?php
fscanf(STDIN, "%d %d", $a, $b);
if ($a === $b) {
    echo ("eq") . PHP_EOL;
} else {
    echo max($a, $b) . PHP_EOL;
}
