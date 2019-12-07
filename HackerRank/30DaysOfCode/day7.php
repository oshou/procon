<?php

$n = intval(fgets(STDIN));
$arr = explode(" ", fgets(STDIN));

$reversed = array_reverse($arr);
for ($i = 0; $i < $n; $i++) {
    printf("%d ", $reversed[$i]);
}
printf("\n");
