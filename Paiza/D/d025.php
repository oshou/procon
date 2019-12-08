<?php

$n = intval(fgets(STDIN));
switch (strlen($n)) {
    case 3:
        echo $n;
        break;
    case 2:
        echo "0" . $n;
        break;
    case 1:
        echo "00" . $n;
}
echo "\n";
