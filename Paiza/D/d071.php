<?php

fscanf(STDIN, "%d %d", $t, $m);
$isTarget = (($t >= 25 && $m <= 40) || ($t < 25 && $m < 40));
if ($isTarget) {
    print("No\n");
} else {
    print("Yes\n");
}
