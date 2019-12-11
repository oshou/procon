<?php
$s = trim(fgets(STDIN));
$map = [
    0 => "C",
    1 => "A",
    2 => "B",
];
for ($i = 0; $i < strlen($s); $i++) {
    printf("%s", $map[$s[$i]]);
}
print("\n");
