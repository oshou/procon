<?php
$n = intval(trim(fgets(STDIN)));
$sum = 0;
for ($i = 1; $i <= $n; $i++) {
    $sum += $i;
}
echo $sum . PHP_EOL;
