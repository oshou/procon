<?php
$s1 = trim(fgets(STDIN));
$s2 = trim(fgets(STDIN));
if (strlen($s1) === strlen($s2)) {
    echo "Yes" . PHP_EOL;
} else {
    echo "No" . PHP_EOL;
}
