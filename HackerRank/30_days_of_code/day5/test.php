<?php

$nums = [1, 2, 3, 4, 5];
for ($i = 0; $i < 5; $i++) {
    printf("index=%d, value=%d\n", $i, $nums[$i]);
}

foreach ($nums as $num) {
    printf("value=%d\n", $num);
}

foreach ($nums as $key => $value) {
    printf("index=%d, value=%d\n", $key, $value);
}

$people = [
    "name" =>  "sato",
    "age" => 18,
    "sex" => "male",
];

foreach ($people as $key => $value) {
    printf("index=%s, value=%s\n", $key, $value);
}

$people = [
    [
        "name" =>  "sato",
        "age" => 18,
        "sex" => "male",
    ],
    [
        "name" =>  "saito",
        "age" => 20,
        "sex" => "female",
    ],
    [
        "name" =>  "kato",
        "age" => 48,
        "sex" => "male",
    ]
];

foreach ($people as $person) {
    foreach ($person as $key => $value) {
        printf("index=%s, value=%s\n", $key, $value);
    }
}

echo ("111" == 111) ? "OK!" : "NG!";    #=> OK!
echo ("111" === 111) ? "OK!" : "NG!";   #=> NG!

$str = "aaa111";
if is_string($str) and is
