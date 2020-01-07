<?php

// Complete the solve function below.
function solve($meal_cost, $tip_percent, $tax_percent)
{
    $total_percent = 100 + $tip_percent + $tax_percent;
    $total_cost = round($meal_cost * $total_percent / 100);

    print ($total_cost) . PHP_EOL;
}

$stdin = fopen("php://stdin", "r");

fscanf($stdin, "%lf\n", $meal_cost);

fscanf($stdin, "%d\n", $tip_percent);

fscanf($stdin, "%d\n", $tax_percent);

solve($meal_cost, $tip_percent, $tax_percent);

fclose($stdin);
