<?php
$n = intval(trim(fgets(STDIN)));
$str = "";
for ($i = 0; $i < $n; $i++) {
    $str .= trim(fgets(STDIN));
}
echo $str . PHP_EOL;
