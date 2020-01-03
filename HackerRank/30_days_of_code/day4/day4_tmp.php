<?php

class Person
{
    public $age;

    public function __construct($initialAge)
    {
        if ($initialAge < 0) {
            print ("Age is not valid, setting age to 0.") . PHP_EOL;
            $this->age = 0;
        } else {
            $this->age = $initialAge;
        }
    }

    public function yearPasses(): void
    {
        $this->age++;
    }

    public function amIOld(): void
    {
        if ($this->age < 13) {
            print ("You are young.") . PHP_EOL;
        } elseif ($this->age < 18) {
            print ("You are a teenager.") . PHP_EOL;
        } else {
            print ("You are old.") . PHP_EOL;
        }
    }
}

$T = intval(fgets(STDIN));
$ages = [];
for ($i = 1; $i <= $T; $i++) {
    $ages[] = intval(fgets(STDIN));
}

foreach ($ages as $age) {
    $person = new Person($age);
    $person->amIOld();
    $person->yearPasses();
    $person->yearPasses();
    $person->yearPasses();
    $person->amIOld();
    echo PHP_EOL;
}
