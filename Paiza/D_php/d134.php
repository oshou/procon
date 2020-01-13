<?php
$s = trim(fgets(STDIN));
for ($i = 0; $i < strlen($s); $i++) {
    echo $s[$i];
    if (($i + 1) % 10 === 0) {
        echo PHP_EOL;
    }
}
