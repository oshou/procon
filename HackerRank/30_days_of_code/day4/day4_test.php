<?php

require __DIR__ . "/day4.php";

use PHPUnit\Framework\TestCase;

class PersonTest extends TestCase
{
    function testAgeNegative()
    {
        $person = new Person(-1);
        $this->assertSame("You are young.", $person->amIOld());
    }

    function testAgeUnder13()
    {
        $person = new Person(5);
        $this->assertSame("You are young.", $person->amIOld());
        $person = new Person(12);
        $this->assertSame("You are young.", $person->amIOld());
    }

    function testAgeUnder18()
    {
        $person = new Person(13);
        $this->assertSame("You are a teenager.", $person->amIOld());
        $person = new Person(17);
        $this->assertSame("You are a teenager.", $person->amIOld());
    }

    function testAgeOver18()
    {
        $person = new Person(18);
        $this->assertSame("You are old.", $person->amIOld());
    }

    function testAgePasses()
    {
        $person = new Person(12);
        $this->assertSame("You are young.", $person->amIOld());
        $person->yearPasses();
        $this->assertSame("You are a teenager.", $person->amIOld());
        $person->yearPasses();
        $person->yearPasses();
        $person->yearPasses();
        $person->yearPasses();
        $this->assertSame("You are a teenager.", $person->amIOld());
        $person->yearPasses();
        $this->assertSame("You are old.", $person->amIOld());
    }
}
