<?php
$ng = trim(fgets(STDIN));
$text = trim(fgets(STDIN));
if (substr_count($text, $ng) === 0) {
    echo $text . PHP_EOL;
} else {
    echo "NG" . PHP_EOL;
}
