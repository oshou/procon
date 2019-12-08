<?php

$s = trim(fgets(STDIN));
$t = trim(fgets(STDIN));
if ($s === $t) {
    echo "Yes\n";
} else {
    echo "No\n";
}
