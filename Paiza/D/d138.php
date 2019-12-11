<?php
fscanf(STDIN, "%d %d", $w, $h);
for ($i = 0; $i < $h; $i++) {
    printf("%s\n", trim(fgets(STDIN)));
}
