<?php
fscanf(STDIN, "%d %d %d", $a, $b, $c);
echo max($c - $a + $b, 0);
