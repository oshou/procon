<?php

$n = intval(fgets(STDIN));
const PI = "3.141592653589793";
printf("%s\n", substr(PI, "0", 1 + $n));
