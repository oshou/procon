<?php
$s = trim(fgets(STDIN));
$n = intval(fgets(STDIN));
switch ($s) {
    case "S":
        echo 1925 + $n . PHP_EOL;
        break;
    case "H":
        echo 1988 + $n . PHP_EOL;
        break;
}
