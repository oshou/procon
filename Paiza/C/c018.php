<?php
$counts = [];
$n = intval(trim(fgets(STDIN)));
for ($i = 0; $i < $n; $i++) {
    fscanf(STDIN, "%s %d", $a[$i], $b[$i]);
}
$m = intval(trim(fgets(STDIN)));
for ($j = 0; $j < $m; $j++) {
    fscanf(STDIN, "%s %d", $c[$j], $d[$j]);
}

for ($i = 0; $i < $n; $i++) {
    $pos = array_search($a[$i], $c);
    if ($pos === false) {
        continue;
    }
    $counts[$i] = floor($d[$pos] / $b[$i]);
}
echo min($counts) ?: 0;
