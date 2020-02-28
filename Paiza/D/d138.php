<?php
fscanf(STDIN, "%d %d", $w, $h);
for ($i = 0; $i < $h; $i++) {
    echo trim(fgets(STDIN)) . PHP_EOL;
}
