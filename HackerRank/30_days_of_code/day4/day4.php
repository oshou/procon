<?php

class Person
{
    public $age;
    public function __construct($initialAge)
    {
        if ($initialAge < 0) {
            echo "Age is not valid, setting age to 0." . PHP_EOL;
            $this->age = 0;
        }
        $this->age = $initialAge;
    }

    public function amIOld(): string
    {
        if ($this->age < 13) {
            return "You are young.";
        }

        if ($this->age < 18) {
            return "You are a teenager.";
        }

        return "You are old.";
    }

    public function yearPasses(): void
    {
        $this->age++;
    }
}
