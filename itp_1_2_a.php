<?php
fscanf(STDIN, "%d %d", $a, $b);
if ($a < $b){
    echo "a < b";
} elseif ($a > $b){
    echo "a > b";
} else {
    echo "a = b";
}
