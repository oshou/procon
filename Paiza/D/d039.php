<?php

$a = intval(fgets(STDIN));
$b = intval(fgets(STDIN));
$c = intval(fgets(STDIN));

if ($a === $b && $b === $c) {
    echo "YES\n";
} else {
    echo "NO\n";
}
