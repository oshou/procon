<?php

$ng_word = trim(fgets(STDIN));
$word = trim(fgets(STDIN));
if (!strpos($word, $ng_word)) {
    printf("%s\n", $word);
} else {
    print("NG\n");
}
