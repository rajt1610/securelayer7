<?php

/*
5) Write a program to accept URL slugs with only alphabets, hyphens and numbers with specific format explained in constraints. Apart from alphabets, numbers and hyphens no other characters are allowed.
Constraints:
(a) Use regex
(b) Numbers are allowed only at end of the strings in URLs.

E.g. Allowed: something4-name, something-name4, something4-name4, something-name
Rejected: 4something-name, some4thing-name, something-4name, etc.

Input:
something-name

Output:
1 (accepted)

Input:
something4-##@name

Output:
0 (rejected)

*/
$input = "something-name";


$pattern = "/^[a-zA-Z]+[0-9]*(-[a-zA-Z]+[0-9]*)*$/";

if (preg_match($pattern, $input)) {
    echo 1;
} else {
    echo 0; 
}
?>