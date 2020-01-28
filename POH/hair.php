<?php

$yes = 0;
$no = 0;
for ($i = 0; $i < 5; $i++) {
    $s = trim(fgets(STDIN));
    switch ($s) {
        case "yes":
            $yes++;
            break;
        case "no":
            $no++;
            break;
    }
}
if ($yes > $no) {
    echo "yes" . PHP_EOL;
} else {
    echo "no" . PHP_EOL;
}
