We are given a list of valid pins ["1234", "0001", "6789"]
We are also given an input string from a pin pad. The input string contains the numbers [0-9] and can contain backspace characters represented by '#'.
Write a function to check if a string is a valid pin number.
Input pin: "1234" -> true
Input pin: "0000" -> false
Input pin: "12344#" -> true it resolves to "1234"
Input pin: "02#0011#" -> true it resolves to "0001"