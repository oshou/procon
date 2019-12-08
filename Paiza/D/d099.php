<?php
$str = trim(fgets(STDIN));
for ($i = 0; $i < strlen($str); $i++) {
    printf("%s\n", $str[$i]);
}
