<?php
fscanf(STDIN, "%d %d", $m, $n);
$collect = 0;
$remain = $n;
$product = $n;
for ($i = 0; $i < $m; $i++) {
    $num = intval(fgets(STDIN));
    $p = intdiv($n, $num);
    $r = $n % $num;
    var_dump($p);
    var_dump($r);
    if ((is_null($remain)) || ($r < $remain)) {
        $remain = $r;
        $collect = $i;
        continue;
    }
    if (($r == $remain)) {
        if ((is_null($product)) || ($p >= $product)) {
            $procut = $p;
            $collect = $i;
            continue;
        }
    }
}
printf("%d\n", $collect);
