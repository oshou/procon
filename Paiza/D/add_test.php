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
            [4, 2, 2],
            [14, 6, 8],
            [44, 11, 33],
            [35, 12, 23],
        ];
    }
}
