<?php
fscanf(STDIN, "%s %s", $s1, $s2);
if ($s2 === "M") {
    echo "Hi, Mr. " . $s1 . "\n";
} else {
    echo "Hi, Ms. " . $s1 . "\n";
}
