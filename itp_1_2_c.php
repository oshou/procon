<?php

fscanf(STDIN,"%d %d %d", $a, $b, $c);

$arr = array($a, $b, $c);
sort($arr);
printf("%d %d %d\n", $arr[0], $arr[1], $arr[2]);
