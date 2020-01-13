<?php

require __DIR__ . "/add.php";

use PHPUnit\Framework\TestCase;

class AddTest extends TestCase
{
    /**
     * @dataProvider addProvider
     */
    public function testAdd($expected, $a, $b)
    {
        $this->assertSame($expected, add($a, $b));
    }

    public function addProvider()
    {
        return [
            [0, 0, 0],
            [3, 1, 2],
        ];
    }
}
