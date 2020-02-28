<?php
$n = intval(trim(fgets(STDIN)));
$s = trim(fgets(STDIN));
printf("%d %d\n", substr_count($s, "S"), substr_count($s, "R"));
