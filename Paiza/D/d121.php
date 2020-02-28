<?php
$s = trim(fgets(STDIN));
if ($s[0] === "A") {
    $s[0] = "R";
}
echo $s . PHP_EOL;
