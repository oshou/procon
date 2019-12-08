<?php
function pricePerArea(int $x, $y, $p)
{
    return $p / ($x * $y);
}

fscanf(STDIN, "%d %d %d", $x1, $y1, $p1);
fscanf(STDIN, "%d %d %d", $x2, $y2, $p2);
$pricePerArea1 = pricePerArea($x1, $y1, $p1);
$pricePerArea2 = pricePerArea($x2, $y2, $p2);
if ($pricePerArea1 === $pricePerArea2) {
    print("DRAW\n");
} elseif ($pricePerArea1 > $pricePerArea2) {
    printf("%d %d %d", $x2, $y2, $p2);
} elseif ($pricePerArea1 < $pricePerArea2) {
    printf("%d %d %d", $x1, $y1, $p1);
}
