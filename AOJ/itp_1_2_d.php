<?php
fscanf(STDIN,"%d %d %d %d %d", $W, $H, $x, $y, $r);

if ( $x >= $r && $y >= $r && ($x + $r) <= $W && ($y + $r) <= $H){
    echo 'Yes';
} else {
    echo 'No';
}

echo PHP_EOL;
