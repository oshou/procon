<?php
fscanf(STDIN, "%d %d %d", $x1, $y1, $p1);
fscanf(STDIN, "%d %d %d", $x2, $y2, $p2);
$r1 = ($p1 / ($x1 * $y1));
$r2 = ($p2 / ($x2 * $y2));
if ($r1 === $r2) {
    print("DRAW\n");
} else if ($r1 > $r2) {
    printf("%d %d %d\n", $x2, $y2, $p2);
} else {
    printf("%d %d %d\n", $x1, $y1, $p1);
}
