<?php
fscanf(STDIN, "%d %d", $n, $m);
$modn = $n % 2;
$modm = $m % 2;
if ($modn !== $modm) {
    echo "YES" . PHP_EOL;
} else {
    echo "NO" . PHP_EOL;
}
