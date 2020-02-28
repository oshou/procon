<?php
fscanf(STDIN, "%d %s", $n, $s);
switch ($s) {
    case "km":
        echo $n * 10 * 100 * 1000;
        break;
    case "m":
        echo $n * 10 * 100;
        break;
    case "cm":
        echo $n * 10;
        break;
}
echo PHP_EOL;
