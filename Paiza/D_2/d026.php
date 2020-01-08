<?php
$arr = [];
for ($i = 0; $i < 7; $i++) {
    $arr[] = trim(fgets(STDIN));
}

$cnt = 0;
foreach ($arr as $value) {
    if ($value === "no") {
        $cnt++;
    }
}
echo $cnt . PHP_EOL;
