<?php
fscanf(STDIN, "%s %d", $s, $n);
for ($i = 0; $i < strlen($s); $i++) {
    if ($i === $n - 1) {
        echo $s[$i] . PHP_EOL;
        exit;
    }
}
