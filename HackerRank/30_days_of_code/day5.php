<?php

$stdin = fopen("php://stdin", "r");
fscanf($stdin, "%d\n", $n);

const MAX = 10;
for ($i = 1; $i <= MAX; $i++) {
    printf("%d x %d = %d\n", $n, $i, $n * $i);
}

fclose($stdin);
