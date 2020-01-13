<?php
$s = trim(fgets(STDIN));
$ans = "unknown";
switch ($s[0]) {
    case "2":
        $ans = "OK";
        break;
    case "4":
        $ans = "error";
        break;
}
echo $ans . PHP_EOL;
