<?php
$n = intval(fgets(STDIN));
$s = trim(fgets(STDIN));
printf("%s\n", str_repeat($s, $n));
