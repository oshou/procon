<?php
$n = intval(fgets(STDIN));
$arr = [];
for ($i = 0; $i < $n; $i++) {
    $arr[] = trim(fgets(STDIN));
}

echo "Hello ";
for ($i = 0; $i < $n; $i++) {
    if ($i >= 1) {
        echo ",";
    }
    echo $arr[$i];
}
echo ".";
