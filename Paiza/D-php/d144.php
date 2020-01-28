<?php
$s = trim(fgets(STDIN));
$ans = 0;
for ($i = 0; $i < strlen($s); $i++) {
    $ans += $s[-1 - $i] * 2 ** $i;
}
echo $ans . PHP_EOL;
