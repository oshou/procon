<?php
fscanf(STDIN, "%s %s", $c1, $c2);
if ($c1 === "J" && $c2 === "J") {
    print("J Q");
} else {
    printf("%s %s\n", $c1, $c2);
}
