<?php
$s = trim(fgets(STDIN));
$reversed = "";
for ($i = strlen($s) - 1; $i >= 0; $i--) {
    $reversed .= $s[$i];
}
echo $reversed . PHP_EOL;
