<?php
$s = trim(fgets(STDIN));
$numString = substr($s, 1);
if ($numString === str_repeat($numString[0], strlen($numString))) {
    print("Yes\n");
} else {
    print("No\n");
}
