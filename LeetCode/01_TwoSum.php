<?php

class Solution
{
    function twoSum($nums, $target)
    {
        for ($i = 0; $i < count($nums); $i++) {
            for ($j = $i + 1; $j < count($nums); $j++) {
                if ($nums[$i] + $nums[$j] == $target) {
                    return [$i, $j];
                }
            }
        }
    }
}

$nums = [1, 4, 8, 3, 2, 9, 15];
$target = 13;

$sol = new Solution;
print_r($sol->twoSum($nums, $target));
