<?php
$s = trim(fgets(STDIN));
$ans = "";
foreach (str_split($s) as $c) {
    if (is_numeric($c)) {
        $ans .= $c;
    }
}
echo $ans . PHP_EOL;
