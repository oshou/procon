<?php
fscanf(STDIN, "%d %d %d", $x, $y, $p);
$num = ceil($x / $y);
printf("%d\n", $p * $num);
