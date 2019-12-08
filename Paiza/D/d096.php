<?php
$s = trim(fgets(STDIN));
$pattern = "/I|l|i/";
if (preg_match($pattern, $s)) {
    print("caution\n");
} else {
    printf("%s\n", $s);
}
