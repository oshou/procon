<?php
fscanf(STDIN, "%d %d %d %d", $a, $b, $c, $d);
printf("%d\n", abs($a * $d - $b * $c) / 2);
