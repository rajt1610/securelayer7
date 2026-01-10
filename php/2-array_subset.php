<?php
/*  2) Write a program to check if one array is subset of another another array.
Constraints:
(a) Input 1 is always going to be a single dimensional array
(b) Input 2 can be a multi dimensional array with no limit on depth
(c) Output values: 1 (if it is a subset), 0 (if it is not a subset)

Example:
Input 1 (single dimensional array) :
array('e' => 'dragonfruit', 'f' => 'mango')
Input 2 (multi dimensional array):
$a = array(
array('a' => 'apple', 'b' => 'banana'),
array('c' => 'dragonfruit', array(
array('e' => 'dragonfruit', 'f' => 'mango')
)),
array('g' => Guava, 'h' => 'Avocado')
)

Output:
1

*/
$input1 = array('e' => 'dragonfruit', 'f' => 'mango');


$input2 = array(
    array('a' => 'apple', 'b' => 'banana'),
    array('c' => 'dragonfruit', array(
        array('e' => 'dragonfruit', 'f' => 'mango')
    )),
    array('g' => 'Guava', 'h' => 'Avocado')
);

$allValues = [];

function collectValues($arr, &$allValues) {
    foreach ($arr as $item) {
        if (is_array($item)) {
           
            collectValues($item, $allValues);
        } else {
          
            $allValues[] = $item;
        }
    }
}

collectValues($input2, $allValues);

$isSubset = 1; 
foreach ($input1 as $value) {
    if (!in_array($value, $allValues, true)) {
        $isSubset = 0;
        break;
    }
}


echo $isSubset;  
