<?php
fscanf(STDIN, "%d %d %d", $n, $m, $l);
$c = explode(" ", trim(fgets(STDIN)));
$x = [];
for ($j = 0; $j < $m; $j++) {
    $x[$j] = explode(" ", trim(fgets(STDIN)));
}

// 配列作成
for ($j = 0; $j < $m; $j++) {
    $s[$j] = 0;
    for ($k = 0; $k < $n; $k++) {
        $s[$j] += $c[$k] * $x[$j][$k];
    }
    $s[$j] = round($s[$j], 0);
}
rsort($s);

// 出力
for ($i = 0; $i < $l; $i++) {
    echo $s[$i] . PHP_EOL;
}
