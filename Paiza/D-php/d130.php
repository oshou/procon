<?php
fscanf(STDIN, "%d %d %d %d", $a, $b, $c, $d);
echo floor(abs($a * $d - $b * $c) / 2) . PHP_EOL;
