<?php

fscanf(STDIN, '%d', $n);
$line = trim(fgets(STDIN));
$arr = explode(' ', $line);
$arr = array_reverse($arr);
echo implode(' ',$arr), PHP_EOL;
