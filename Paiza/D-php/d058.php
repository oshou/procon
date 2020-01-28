<?php
fscanf(STDIN, "%d %d %d", $l, $m, $n);
printf("%s%s%s\n", str_repeat("A", $l), str_repeat("B", $m), str_repeat("A", $n));
