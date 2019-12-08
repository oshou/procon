<?php
$arr = explode(" ", trim(fgets(STDIN)));
$sum = 0;
for ($i = 0; $i < count($arr); $i++) {
    if ($arr[$i] > 5) {
        $sum += 5;
    } else {
        $sum += $arr[$i];
    }
}
echo $sum . "\n";
