<?php
fscanf(STDIN, "%s %s", $first, $last);
$InitialF = strtoupper($first[0]);
$InitialL = strtoupper($last[0]);
echo $InitialF . "." . $InitialL . "\n";
