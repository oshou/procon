<?php

$s = trim(fgets(STDIN));
if ($s[0] === "2") {
    print("ok\n");
} elseif ($s[0] === "4") {
    print("error\n");
} else {
    print("unknown\n");
}
