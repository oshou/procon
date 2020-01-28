<?php
$s = trim(fgets(STDIN));
switch ($s) {
    case "candy":
    case "chocolate":
        echo "Thanks!" . PHP_EOL;
        break;
    default:
        echo "No!" . PHP_EOL;
}
