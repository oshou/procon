<?php
$nums = explode(" ", trim(fgets(STDIN)));
$avg = (string) array_sum($nums);
echo $avg[-1] . PHP_EOL;
