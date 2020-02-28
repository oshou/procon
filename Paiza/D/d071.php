<?php
fscanf(STDIN, "%d %d", $t, $m);
$dryNG = ($t < 25 && $m > 40) || ($t >= 25 && $m <= 40);
echo ($dryNG) ? "No" : "Yes";
echo PHP_EOL;
