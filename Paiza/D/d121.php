<?php
$s = trim(fgets(STDIN));
if ($s[0] === "A") {
    $s[0] = "R";
}
printf("%s\n", $s);
