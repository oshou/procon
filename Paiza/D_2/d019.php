<?php
fscanf(STDIN, "%s %d", $s, $n);
echo substr($s, $n - 1, 1) . PHP_EOL;
