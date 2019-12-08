<?php

$n = intval(fgets(STDIN));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[] = trim(fgets(STDIN));
}
printf("%s\n", implode("", $arr));
