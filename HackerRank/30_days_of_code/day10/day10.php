<?php

$dec = intval(fgets(STDIN));
$binString = decbin($dec);

preg_match_all("/(.)\\1*/", $binString, $elements);
$maxlen = 0;
foreach ($elements as $element) {
    $len = max(array_map('strlen', $element));
    if ($len > $maxlen) {
        $maxlen = $len;
    }
}
print ($maxlen) . PHP_EOL;
