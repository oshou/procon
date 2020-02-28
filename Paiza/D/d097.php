<?php
$days = explode(" ", trim(fgets(STDIN)));
echo (array_sum($days) >= 5) ? "yes" : "no";
echo PHP_EOL;
