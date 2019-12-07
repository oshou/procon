<?php

function calculation(int $N): string
{
    if ($N % 2 !== 0) {
        return "Weird";
    }

    if (2 <= $N && $N <= 5) {
        return "Not Weird";
    }

    if (6 <= $N && $N <= 20) {
        return "Weird";
    }

    if (20 <= $N) {
        return "Not Weird";
    }
}

$stdin = fopen("php://stdin", "r");

fscanf($stdin, "%d\n", $N);

print (calculation($N)) . PHP_EOL;

fclose($stdin);
