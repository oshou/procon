<?php
$s = trim(fgets(STDIN));
switch ($s) {
    case "chocolate":
    case "candy":
        echo "Thanks!" . PHP_EOL;
        break;
    default:
        echo "No!" . PHP_EOL;
}
