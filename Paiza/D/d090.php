<?php
$arr = explode(" ", trim(fgets(STDIN)));
$sumString = (string) array_sum($arr);
printf("%s\n", $sumString[-1]);
