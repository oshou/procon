<?php
$arr = explode(" ", trim(fgets(STDIN)));
$sum = 0;
foreach ($arr as $ele) {
    $ele = intval($ele);
    if ($ele >= 5) {
        $sum += 5;
    } else {
        $sum += $ele;
    }
}
echo $sum . PHP_EOL;
