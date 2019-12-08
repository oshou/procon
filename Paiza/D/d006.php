<?php

fscanf(STDIN, "%d %s", $n, $s);

switch ($s) {
    case "km":
        echo $n * 1000 * 100 * 10;
        break;
    case "m":
        echo $n * 100 * 10;
        break;
    case "cm":
        echo $n * 10;
}
echo PHP_EOL;
