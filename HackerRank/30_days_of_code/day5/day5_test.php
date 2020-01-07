<?php

require __DIR__ . "/day5.php";

use PHPUnit\Framework\TestCase;

class PhoneBookTest extends TestCase
{
    function testGetPhoneNumber()
    {
        $phoneBook = new PhoneBook("sam", "99912222");
        $phoneNumber = "99912222";
        $name = "sam";
        $this->assertSame($phoneNumber, $phoneBook->getPhoneNumber($name));
    }
}
