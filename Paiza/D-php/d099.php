<?php
$s = trim(fgets(STDIN));
for ($i = 0; $i < strlen($s); $i++) {
    echo $s[$i] . PHP_EOL;
}
