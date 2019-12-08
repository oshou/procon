<?php
$s = trim(fgets(STDIN));
$underscore = substr_count($s, "_");
$hyphen = substr_count($s, "-");

if ($underscore >= $hyphen) {
    $fromOperator = "-";
    $toOperator = "_";
} else {
    $fromOperator = "_";
    $toOperator = "-";
}
printf("%s\n", str_replace($fromOperator, $toOperator, $s));
