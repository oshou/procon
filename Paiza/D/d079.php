<?php
$s = fgets(STDIN);
$same_str = str_repeat($s[0], strlen($s));
if ($s === $same_str) {
    print("OK\n");
} else {
    print("NG\n");
}
