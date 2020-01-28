<?php
$s = trim(fgets(STDIN));
if (strlen($s) == substr_count($s, $s[0])) {
    echo "NG" . PHP_EOL;
} else {
    echo "OK" . PHP_EOL;
}
