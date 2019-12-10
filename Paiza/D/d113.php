<?php

fscanf(STDIN, "%d:%d", $h, $m);
if ($h - 8 >= 0) {
    printf("%d:%d", $h - 8, $m);
} else {
    printf("%d:%d", $h - 8 + 24, $m);
}
