<?php
$n = intval(fgets(STDIN));
$arr = explode(" ", trim(fgets(STDIN)));
echo $arr[$n - 1] . "\n";
