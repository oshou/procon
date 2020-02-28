<?php
$s1 = trim(fgets(STDIN));
$s2 = trim(fgets(STDIN));
if ($s1[-1] === $s2[0] && $s2[-1] !== "n") {
    echo "OK" . PHP_EOL;
} else {
    echo "NG" . PHP_EOL;
}
