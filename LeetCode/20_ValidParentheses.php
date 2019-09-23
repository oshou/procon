<?php

class Solution
{
    function is_valid($s)
    {
        $stack = [];
        $match = ["(" => ")", "{" => "}", "[" => "]"];
        for ($i = 0; $i < strlen($s); $i++) {
            print_r($stack);
            print_r("cccc" . PHP_EOL);
            print_r("match: " . $match[$s[$i]] . PHP_EOL);
            if (in_array($s[$i], ["(", "{", "["])) {
                array_push($stack, $s[$i]);
            } else {
                echo "else" . PHP_EOL;
                if (isset($stack)) {
                    echo "else-isset" . PHP_EOL;
                    return false;
                } else if (end($stack) == $match[$s[$i]]) {
                    echo "else-end-equall" . PHP_EOL;
                    array_pop($stack);
                } else {
                    echo "else-else" . PHP_EOL;
                    return false;
                }
            }
        }
        return (isset($stack));
    }
}

$sol = new Solution;
$s = "({[]})";
echo ($sol->is_valid($s)) . PHP_EOL;
