<?php
$n = intval(fgets(STDIN));
$arr = explode(" ", trim(fgets(STDIN)));
$gCount = count(array_keys($arr, "G"));
$pCount = count(array_keys($arr, "P"));
if ($gCount == $pCount) {
    print("Draw\n");
} elseif ($gCount > $pCount) {
    print("P\n");
} else {
    print("G\n");
}
