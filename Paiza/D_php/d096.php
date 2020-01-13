<?php
$s = trim(fgets(STDIN));
$pattern = '/[Iil]/';
if (preg_match($pattern, $s) === 0) {
    echo $s . PHP_EOL;
} else {
    echo "caution" . PHP_EOL;
}
