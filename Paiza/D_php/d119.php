<?php
$n = intval(trim(fgets(STDIN)));
$pi = "3.141592653589793";
printf("3.%d", substr($pi, 2, $n));
