<?php

$n = (int)trim(fgets(STDIN));
$line = trim(fgets(STDIN));
$arr = explode(" ", $line);
echo implode(" ", array_reverse($arr)), PHP_EOL;
