<?php
$n = intval(fgets(STDIN));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[] = trim(fgets(STDIN));
}
printf("Hello %s.\n", implode(",", $arr));
