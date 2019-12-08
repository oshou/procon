<?php
$arr = explode(" ", trim(fgets(STDIN)));
$cnt = 0;
foreach ($arr as $ele) {
    if ($ele <= 2) {
        $cnt++;
    }
}
printf("%d\n", $cnt);
