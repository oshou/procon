<?php
$leet = [
    "A" => "4",
    "E" => "3",
    "G" => "6",
    "I" => "1",
    "O" => "0",
    "S" => "5",
    "Z" => "2",
];
$leetKeys = ["A", "E", "G", "I", "O", "S", "Z"];
$s = trim(fgets(STDIN));
for ($i = 0; $i < strlen($s); $i++) {
    if (in_array($s[$i], $leetKeys)) {
        $s[$i] = $leet[$s[$i]];
    }
}
echo $s . PHP_EOL;
