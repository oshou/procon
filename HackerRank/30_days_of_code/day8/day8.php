<?php

function getPhoneNumber(array $arr, string $name)
{
    $index = array_search($name, array_column($arr, 'name'));
    if (!is_int($index)) {
        return "Not found";
    }
    $phoneNumber = $arr[$index]["phone_number"];
    return sprintf("%s=%d", $name, $phoneNumber);
}


// テスト数入力
$n = intval(fgets(STDIN));
// テストデータ入力
$arr = [];
for ($i = 0; $i < $n; $i++) {
    fscanf(
        STDIN,
        "%s %d¥n",
        $arr[$i]["name"],
        $arr[$i]["phone_number"]
    );
}
// ユーザー名指定で番号出力
for ($i = 0; $i < $n; $i++) {
    fscanf(STDIN, "%s", $name);
    print (getPhoneNumber($arr, $name)) . PHP_EOL;
}
