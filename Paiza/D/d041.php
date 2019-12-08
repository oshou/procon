<?php
fscanf(STDIN, "%d %d %d", $n, $d, $e);
if ($n <= $d * $e) {
    echo "OK\n";
} else {
    echo "NG\n";
}
