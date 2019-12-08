<?php

$cnt = 0;
for ($i = 0; $i < 7; $i++) {
    $percent = intval(fgets(STDIN));
    if ($percent <= 30) {
        $cnt++;
    }
}
echo $cnt . "\n";
