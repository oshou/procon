<?php
$map = [
    0 => "C",
    1 => "A",
    2 => "B",
];
$s = trim(fgets(STDIN));
$txt = "";
for ($i = 0; $i < strlen($s); $i++) {
    $txt .= $map[$s[$i]];
}
echo $txt . PHP_EOL;
