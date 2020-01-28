<?php
$s = trim(fgets(STDIN));
echo preg_replace('/False/', 'True', $s) . PHP_EOL;
