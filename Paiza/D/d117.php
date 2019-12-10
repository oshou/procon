<?php
$n = intval(fgets(STDIN));
$arr = explode(" ", fgets(STDIN));
printf("%d\n", array_sum($arr));
