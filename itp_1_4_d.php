<?php

fscanf(STDIN,'%d',$n);
$line = trim(fgets(STDIN));
$arr = explode(' ',$line);
printf('%d %d %d%s', min($arr), max($arr), array_sum($arr), PHP_EOL);
