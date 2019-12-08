<?php
$a = trim(fgets(STDIN));
$b = trim(fgets(STDIN));
if ($a[-1] === $b[0] && $b[-1] !== "n") {
    print("OK\n");
} else {
    print("NG\n");
}
