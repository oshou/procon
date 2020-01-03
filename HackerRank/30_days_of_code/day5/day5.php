<?php

class Phonebook
{
    private $phonebook;

    function __construct($name, $phoneNumber)
    {
        $this->phonebook[] = [
            "name" => $name,
            "phone_number" => $phoneNumber,
        ];
    }

    function getPhoneNumber($name)
    {
        foreach ($this->phonebook as $address) {
            if ($address["name"] === $name) {
                return $address["phone_number"];
            }
        }
    }
}
