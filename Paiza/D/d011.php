<?php
const ALFABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZA";
$c = trim(fgets(STDIN));
echo (strpos(ALFABET, $c) + 1) . PHP_EOL;
