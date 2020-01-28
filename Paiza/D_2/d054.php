<?php
$s = trim(fgets(STDIN));
$cnt = substr_count($s, "1");
if ($cnt >= 11) {
    echo "OK" . PHP_EOL;
} else {
    echo (11 - $cnt) . PHP_EOL;
}
