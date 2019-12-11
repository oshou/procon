<?php
$n = intval(fgets(STDIN));
$a = intval(fgets(STDIN));
$b = intval(fgets(STDIN));
$unit = floor($n / $a);
printf("%d\n", ($b - $a) * $unit);
