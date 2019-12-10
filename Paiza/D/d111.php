<?php

$n = intval(fgets(STDIN));
$s = trim(fgets(STDIN));
printf("%s\n", substr($s, 0, $n));
