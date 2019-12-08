<?php

$n = intval(fgets(STDIN));
if ($n <= 30) {
    echo "sunny\n";
} elseif ($n <= 70) {
    echo "cloudy\n";
} else {
    echo "rainy\n";
}
