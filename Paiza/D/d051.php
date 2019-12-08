<?php

$arr = explode(" ", trim(fgets(STDIN)));
$cnt = 0;
foreach ($arr as $ele) {
    if ($ele === "W") {
        $cnt++;
    }
}
if ($cnt >= 5) {
    echo "OK\n";
} else {
    echo "NG\n";
}
