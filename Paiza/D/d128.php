<?php
$n = intval(fgets(STDIN));
$arr = explode(" ", trim(fgets(STDIN)));
foreach ($arr as $ele) {
    printf("%s", $ele[0]);
}
print("\n");
