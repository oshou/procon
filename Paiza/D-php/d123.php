<?php
$price = intval(fgets(STDIN));
if ($price < 10000) {
    $price += 10000;
}
echo $price . PHP_EOL;
