<?php



/*

1) Write a program to compare two multi dimensional arrays to search for values and if matched count the number of matched items. Make sure each value must exist atleast 2 times. Make a list of those items and create a new array with the value as keys and its count as values.

Example:
Input:
$a = array(
array('a' => 'apple', 'b' => 'banana'),
array('c' => 'dragonfruit', 'd' => 'banana'),
array('e' => 'dragonfruit', 'f' => 'mango')
);
Output:
$result = array('banana'=>2,'dragonfruit'=>2)

*/


$a = array(
    array('a' => 'apple',       'b' => 'banana'),
    array('c' => 'dragonfruit', 'd' => 'banana'),
    array('e' => 'dragonfruit', 'f' => 'mango')
);


$flat = [];
foreach ($a as $row) {
    foreach ($row as $v) {
        $flat[] = $v;
    }
}


$counts = array_count_values($flat);


$result = [];
foreach ($counts as $value => $count) {
    if ($count >= 2) {
        $result[$value] = $count;   
    }
}


$matchedItems = array_keys($result);

echo "Matched Items:\n";
print_r($matchedItems);

echo "Result (value => count):\n";
print_r($result);

/*
Matched Items:
Array ( [0] => banana [1] => dragonfruit )

Result:
Array ( [banana] => 2 [dragonfruit] => 2 )
*/

