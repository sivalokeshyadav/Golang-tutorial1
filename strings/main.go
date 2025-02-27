package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Strings:-It is a sequence of variable-width characters where every character is represented by one or more bytes using UTF-8 Encoding.
// In other words, strings are the immutable chain of arbitrary bytes(including bytes with zero value) or
//
//	string is a read-only slice of bytes and the bytes of the strings can be represented in the Unicode
//
// text using UTF-8 encoding. Due to UTF-8 encoding Golang string can contain a text which is the mixture
//
//	of any language present in the world, without any confusion and limitation of the page.
//
// Generally, strings are enclosed in double-quotes””
// strings cant be empty but they are nil
// string literals:-theyare2types
// 1.usingdoublequotes:-the string literals are created using double-quotes(“”).
// This type of string support escape character but does not span multiple lines.
//
//	This type of string literals is widely used in Golang programs.
//
// escape character:-\",\a,\b,\\,\000,\',\f,\n,\r,\t,\uhhhh,\v,\xhh,"
// 2.usingbackticks:-the string literals are created using backticks(“)
//
//	and also known as raw literals. Raw literals do not support escape
//
// characters, can span multiple lines, and may contain any character
// except backtick. It is, generally, used for writing multiple line message,
//
//	in the regular expressions, and in HTML.
//
// Trim:-The Trim function removes all specified leading and trailing characters from a string.
// syntax:-func Trim(s string, cutset string) string,s:stringtotrim,cutset:-characterstoremovefrombothends
// Trimleft:-removesleftendcharactersfromspecified
// TrimRight:-removesrightendcharactersfromspecified,
// TrimSpace:removes all leading and trailing whitespace from a string.syntax:-func TrimSpace(s string) string
// The TrimPrefix function removes a specified prefix from the beginning of a string if present.
// syntax:-func TrimPrefix(s, prefix string) string
// TrimSuffix function removes a specified suffix from the end of a string if present.
// syntax:-func TrimSuffix(s, suffix string) string
// Contains function checks if a string contains a specified substring.
// syntax:-func Contains(s, substr string) bool
// ContainsRune function checks if a string contains a specified Unicode code point.
// split a string in a golang
// The Split function divides a string into all substrings separated by the specified separator and
// returns a slice with these substrings.
// func Split(s, sep string) []string                             #Using Split Function
// The SplitAfter function splits a string after each instance of the specified separator and returns a slice.
// func SplitAfter(s, sep string) []string                     # Using SplitAfter Function
// The SplitN function splits a string into a maximum number of substrings.
// func SplitN(s, sep string, n int) []string                 # Using SplitN Function
// The SplitAfterN function splits a string after each instance of the separator, but only up to n substrings.
// func SplitAfterN(s, sep string, n int) []string       # Using SplitAfterN Function
// compare
// uisng compare operators:-strings can be compared with operators like==,!=,<,>,<=and>=,this allows checking equality or lexicographical ordering
// Using strings.Compare Function
// The strings.Compare function compares two strings lexicographically and returns:
// 0 if s1 is equal to s2
// 1 if s1 is lexicographically greater than s2
// -1 if s1 is lexicographically less than s2
// syntax:-result := strings.Compare(s1, s2)
func main() {
	fmt.Println("Hello, strings") // Printing "Hello, 世界!" in Chinese
	//using shorthand notation
	s := "Jumping japang thumping thapang gili gili ya"
	fmt.Println(s)
	//using var keyword
	var myVariable string
	fmt.Println(myVariable)
	myVariable = "katam tata see you bye bye"
	fmt.Println(myVariable)
	//stringliterarls
	My_value_1 := "Welcome to GeeksforGeeks"

	// Adding escape character
	My_value_2 := "Welcome!\nGeeksforGeeks"

	// Using backticks
	My_value_3 := `Hello!GeeksforGeeks`

	// Adding escape character
	// in raw literals
	My_value_4 := `Hello!\nGeeksforGeeks`

	// Displaying strings
	fmt.Println("String 1: ", My_value_1)
	fmt.Println("String 2: ", My_value_2)
	fmt.Println("String 3: ", My_value_3)
	fmt.Println("String 4: ", My_value_4)
	mystr := "Welcome to siva"

	fmt.Println("String:", mystr)

	/* if you trying to change
	      the value of the string
	      then the compiler will
	      throw an error, i.e,
	     cannot assign to mystr[1]

	   mystr[1]= 'G'
	   fmt.Println("String:", mystr)

	*/

	// String as a range in the for loop
	for index, s := range "GeeksForGeeKs" {

		fmt.Printf("The index number of %c is %d\n", s, index)
	}
	//creating and initializing a string
	str := "Welcome raja"
	// Accessing the bytes of the given string
	for c := 0; c < len(str); c++ {
		fmt.Printf("\nCharacter =%c Bytes=%x", str[c], str[c])
	}
	//create a string from slice
	//createingandinitalizingasliceofbyte
	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	//createingstringfromslice
	mystring1 := string(myslice1)
	fmt.Println("\nString from slice:", mystring1)
	//creatingandinitailizingasliceofrune
	myslice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	// Creating a string from the slice
	mystring2 := string(myslice2)

	// Displaying the string
	fmt.Println("String 2: ", mystring2)
	//lengthofthestring
	mystrlen := "Welcometothegeeksforgeeks????????"
	// Using len() function
	length1 := len(mystrlen)
	// Using RuneCountInString() function
	length2 := utf8.RuneCountInString(mystrlen)
	fmt.Println("string:", mystr)
	fmt.Println("Length 1:", length1)
	fmt.Println("Length 2:", length2)
	//trimstring
	strim := "@@Hello, Geeks!!"
	result1 := strings.Trim(strim, "@!")
	fmt.Println(result1)
	result2 := strings.TrimLeft(strim, "@!")
	fmt.Println(result2)
	result3 := strings.TrimRight(strim, "@!")
	fmt.Println(result3)
	ss := "   Hello, Geeks   "
	result4 := strings.TrimSpace(ss)
	fmt.Println(result4)
	// TrimPrefix function
	prefix := "Hell"
	str1 := "Hello, Geeks"
	result5 := strings.TrimPrefix(str1, prefix)
	fmt.Println(result5)
	// TrimSuffix function
	suffix := "s"
	str2 := "Hello, Geeks"
	result6 := strings.TrimSuffix(str2, suffix)
	fmt.Println(result6)
	//splittingthestring
	str3 := "Welcome, to, Geeks, for, Golang"
	fmt.Println("", str3)
	res := strings.Split(str3, ",")
	fmt.Println(res)

	//splitting the string after occurrence
	result7 := strings.SplitAfter(str3, ",")
	fmt.Println(result7)

	//splitting the string after n occurrence
	res1 := strings.SplitAfterN(str3, ",", 2)
	fmt.Println(res1)
	//splitN
	result8 := strings.SplitN(str3, ",", 2)
	fmt.Println(result8)
	//comparing strings using operators
	s11 := "Hello"
	s21 := "Geeks"
	s31 := "Hello"

	// Equality and inequality
	fmt.Println("s11 == s21:", s11 == s21) // false
	fmt.Println("s11 == s31:", s11 == s31) // true
	fmt.Println("s11 != s21:", s11 != s21) // true

	// Lexicographical comparison
	fmt.Println("s11 < s21:", s11 < s21)   // true
	fmt.Println("s11 > s21:", s11 > s21)   // false
	fmt.Println("s11 <= s31:", s11 <= s31) // true
	fmt.Println("s11 >= s31:", s11 >= s31) // true
	//comparing strings using compare function
	fmt.Println("strings.Compare(s11, s21):", strings.Compare(s11, s21)) // -1
	fmt.Println("strings.Compare(s11, s31):", strings.Compare(s11, s31)) // 0
	fmt.Println("strings.Compare(s21, s11):", strings.Compare(s21, s11)) // 1
}

// strings are immutable once a string is created the value of the string cannot be changed. Or in other words, strings are read-only. If you try to change,
// then the compiler will throw an error.
