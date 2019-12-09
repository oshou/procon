<?php
fscanf(STDIN, "%d %d", $m, $n);
if ($m % 2 != $n % 2) {
    print("YES\n");
} else {
    print("NO\n");
}
