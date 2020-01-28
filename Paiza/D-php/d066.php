<?php
fscanf(STDIN, "%d %d", $m, $n);
$power = $n - $m;
if ($power < 0) {
    echo "No" . PHP_EOL;
} else {
    echo $power . PHP_EOL;
}
