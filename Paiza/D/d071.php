<?php

fscanf(STDIN, "%d %d", $t, $m);
$isTarget = (($t >= 25 && $m >= 40) || ($t < 25 && $m < 40));
if ($isTarget) {
    print("Yes\n");
} else {
    print("No\n");
}
