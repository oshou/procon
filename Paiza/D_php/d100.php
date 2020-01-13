<?php
$s = trim(fgets(STDIN));
$cnt_underscore = substr_count($s, "_");
$cnt_hyphen = substr_count($s, "-");
if ($cnt_hyphen > $cnt_underscore) {
    echo str_replace("_", "-", $s) . PHP_EOL;
} else {
    echo str_replace("-", "_", $s) . PHP_EOL;
}
