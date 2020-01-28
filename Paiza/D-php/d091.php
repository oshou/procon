<?php
$scores = explode(" ", trim(fgets(STDIN)));
$cnt = 0;
foreach ($scores as $score) {
    if ($score <= 2) {
        $cnt++;
    }
}
echo $cnt . PHP_EOL;
