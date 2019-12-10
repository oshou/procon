<?php
fscanf(STDIN, "%s %s", $m, $d);
$str = $m . $d;
if ($str === str_repeat($str[0], strlen($str))) {
    print("Yes\n");
} else {
    print("No\n");
}
