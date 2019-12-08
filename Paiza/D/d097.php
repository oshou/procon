<?php

$arr = explode(" ", trim(fgets(STDIN)));
$value_count = array_count_values($arr);
if ($value_count[1] >= 5) {
    print("yes\n");
} else {
    print("no\n");
}
