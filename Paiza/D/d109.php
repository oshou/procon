<?php
fscanf(STDIN, "%s %s", $m, $d);
$s = $m . $d;
if ($s = str_replace($s[0], strlen($s))) {
    echo "Yes" . PHP_EOL;
} else {
    echo "No" . PHP_EOL;
}
