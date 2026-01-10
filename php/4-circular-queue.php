<?php
/*

4) Write a program to rearrange given string in reverse order. You have to start from the position given as input. The string is to be imagined as a circular queue.
Example:
Input 1:
Chocolate
Input 2:
4
Output:
cohcetalo

*/


$str = "Chocolate";
$start = 4;

$length = strlen($str);
$result = "";

for ($i = 0; $i < $length; $i++) {

 
    $index = ($start - $i) % $length;

    if ($index < 0) {
        $index += $length;
    }

    $result .= strtolower($str[$index]);
}

echo $result;
