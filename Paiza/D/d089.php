<?php
$s = trim(fgets(STDIN));
$str = "";
for ($i = 0; $i < strlen($s); $i++) {
    if (is_numeric($s[$i])) {
        $str .= $s[$i];
    } else {
        break;
    }
}
printf("%s\n", (int) $str);
