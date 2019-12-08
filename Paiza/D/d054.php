<?php

$s = trim(fgets(STDIN));
if (strlen($s) >= 11) {
    echo "OK\n";
} else {
    echo (11 - strlen($s)) . "\n";
}
