<?php
$n = intval(trim(fgets(STDIN)));
$arr = explode(" ", trim(fgets(STDIN)));
$counts = array_count_values($arr);
if ($counts["G"] === $counts["P"]) {
    echo "Draw" . PHP_EOL;
} else if ($counts["G"] > $counts["P"]) {
    echo "P" . PHP_EOL;
} else {
    echo "G" . PHP_EOL;
}
