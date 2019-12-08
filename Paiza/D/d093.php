<?php

$num = (string) intval(fgets(STDIN));
$same_num = str_repeat($num[0], strlen($num));
if ($num == $same_num) {
    printf("%d\n", $num);
} else {
    print("No\n");
}
