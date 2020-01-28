<?php
$n = intval(fgets(STDIN));
$words = explode(" ", trim(fgets(STDIN)));
$str = "";
foreach ($words as $word) {
    $str .= $word[0];
}
echo $str . PHP_EOL;
