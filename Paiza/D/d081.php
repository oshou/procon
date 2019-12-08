<?php
$n = intval(fgets(STDIN));
fscanf(STDIN, "%d %d", $h, $w);
printf("%d\n", ($h * $w) % $n);
