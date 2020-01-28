<?php
$s = trim(fgets(STDIN));
$sub = substr($s, 1, strlen($s) - 1);
if ($sub === str_repeat($sub[0], strlen($sub))) {
    echo "Yes" . PHP_EOL;
} else {
    echo "No" . PHP_EOL;
}
