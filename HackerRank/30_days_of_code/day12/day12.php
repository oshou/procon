<?php

/**
 *
 */
class Person
{
    private $first_name;
    private $last_name;
    private $id;
    private $scores;

    function __construct()
    { }
}

/**
 *
 */
class Student extends Person
{
    function __construct()
    { }

    public function calculate()
    { }
}
