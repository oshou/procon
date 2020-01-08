<?php
$s = trim(fgets(STDIN));
$cnt = 0;
for ($i = 0; $i < strlen($s); $i++) {
    if ($s[$i] === "A") {
        $cnt++;
    }
}
echo $cnt . PHP_EOL;
