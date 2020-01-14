<?php
$endValue = 30;
for ($i = 0; $i < $endValue; $i++) {
    if ($i % 15 === 0) {
        echo "FizzBuzz";
    } else if ($i % 3 === 0) {
        echo "Fizz";
    } else if ($i % 5 === 0) {
        echo "Buzz";
    } else {
        echo $i;
    }
    echo PHP_EOL;
}
