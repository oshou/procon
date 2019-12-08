<?php
$s = trim(fgets(STDIN));
switch ($s) {
    case "candy":
    case "chocolate":
        echo "Thanks!\n";
        break;
    default:
        echo "No!\n";
}
