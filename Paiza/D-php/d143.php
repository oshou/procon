<?php
fscanf(STDIN, "%d %d %d", $m, $v, $f);
echo floor(($m * $v ** 2) / (2 * $f)) . PHP_EOL;
