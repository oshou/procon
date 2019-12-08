<?php

$cnt = 0;
for ($i = 0; $i < 7; $i++) {
    $s = trim(fgets(STDIN));
    if ($s !== "yes") {
        $cnt++;
    }
}
echo $cnt . "\n";
