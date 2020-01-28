<?php
$arr = explode(" ", trim(fgets(STDIN)));
$a = intval(trim(fgets(STDIN)));
$rank = count($arr) + 1;
for ($i = 0; $i < count($arr); $i++) {
    if ($a < $arr[$i]) {
        $rank = $i + 1;
        break;
    }
}
echo $rank . PHP_EOL;
