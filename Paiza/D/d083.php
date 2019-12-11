<?php
fscanf(STDIN, "%d %d", $n, $m);
if ($n + $m < 16) {
    print("HIT\n");
} else {
    print("STAND\n");
}
