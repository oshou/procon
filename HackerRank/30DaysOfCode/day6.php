<?php

$n = intval(fgets(STDIN));
for ($i = 1; $i <= $n; $i++) {
    $input = trim(fgets(STDIN));
    $even = "";
    $odd = "";
    for ($j = 0; $j < strlen($input); $j++) {
        if ($j % 2 === 0) {
            $even .= $input[$j];
        } else {
            $odd .= $input[$j];
        }
    }
    printf("%s %s\n", $even, $odd);
}
