<?php
$g = intval(trim(fgets(STDIN)));
$presents = explode(" ", trim(fgets(STDIN)));
echo $presents[$g - 1] . PHP_EOL;
