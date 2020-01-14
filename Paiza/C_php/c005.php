<?php
$m = intval(trim(fgets(STDIN)));
for ($i = 0; $i < $m; $i++) {
    $s = trim(fgets(STDIN));
    $pattern = '/^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/';
    if (preg_match($pattern, $s) === 0) {
        echo "False" . PHP_EOL;
    } else {
        echo "True" . PHP_EOL;
    }
}
