<?php
$s = trim(fgets(STDIN));
if ($s === str_repeat($s[0], strlen($s))) {
    echo $s . PHP_EOL;
} else {
    echo "No" . PHP_EOL;
}
