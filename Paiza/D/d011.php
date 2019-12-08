<?php

const ALFABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
$s = trim(fgets(STDIN));

echo strpos(ALFABET, $s) + 1 . PHP_EOL;
