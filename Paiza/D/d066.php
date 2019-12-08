<?php
fscanf(STDIN, "%d %d", $m, $n);
if ($n >= $m) {
    printf("%d\n", $n - $m);
} else {
    print("No\n");
}
