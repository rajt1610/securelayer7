


<?php

/*


3) Write a program to find alternate set of 3 elements with highest sum inside an array with size greater than 5.
Constraints:
(a) Elements need to be alternate (e.g. 1st, 3rd, 5th).
(b) After the last element the array ends so it doesnt continue through first elemt again. Its a simple queue which starts from 1st and ends at last.
(c) There is no limit to number of elements in array.
Example:
Input:
$a = array(-1,2,-8,-9,5,-6,5)
Output:
$result = array(-8,5,5)


*/
$a = array(-1, 2, -8, -9, 5, -6, 5);

$maxSum = PHP_INT_MIN;  
$result = [];

$size = count($a);


for ($i = 0; $i + 4 < $size; $i++) {

  
    $currentSet = array(
        $a[$i],
        $a[$i + 2],
        $a[$i + 4]
    );

 
    $sum = $currentSet[0] + $currentSet[1] + $currentSet[2];

 
    if ($sum > $maxSum) {
        $maxSum = $sum;
        $result = $currentSet;
    }
}


print_r($result);
