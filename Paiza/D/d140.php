<?php
$n = trim(fgets(STDIN));
$arr = explode(" ", trim(fgets(STDIN)));
printf("%s\n", $arr[$n - 1]);
