<?php

$s1 = trim(fgets(STDIN));
$s2 = trim(fgets(STDIN));
if (strlen($s1) === strlen($s2)) {
    print("Yes\n");
} else {
    print("No\n");
}
