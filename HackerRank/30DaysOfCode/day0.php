<?php

$_fp = fopen("php://stdin", "r");

$inputString = fgets($_fp);

print("Hello, World.\n");
print($inputString);

fclose($_fp);
