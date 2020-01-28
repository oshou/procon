<?php
fscanf(STDIN, "%d:%d", $h, $m);
if ($h < 8) {
    printf("%d:%d", $h + 16, $m);
} else {
    printf("%d:%d", $h - 8, $m);
}
