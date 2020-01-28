<?php
$arr = explode(" ", trim(fgets(STDIN)));
printf("%.1f\n", array_sum($arr) / count($arr));
