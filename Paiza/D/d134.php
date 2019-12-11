<?php
$s = trim(fgets(STDIN));
for ($i = 0; $i < strlen($s); $i++) {
    printf("%s", $s[$i]);
    if ($i % 10 === 9) {
        print("\n");
    }
}
