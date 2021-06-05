Golang exercises through exercism.io

What I've learned from each exercise

* Accumulate            - ability to pass functions as a parameter to a function
* Anagram               - strings.EqualFold for case insensitive string equality and reflect.DeepEqual for map equality

* Atbash Cipher         - interacting with runes like their ASCII integers (use them in math equations, etc.)

* Bob                   - regexp mustcompile and matchstring, switch/case

* Collatz conjecture    - 

* Darts                 - math sqrt/pow (and hypot)

* Difference of Squares - fast equations for sum of squares and square of sums

* ETL

* Grains               - using bit shifting when needing to double a value, bit shifts are
evaluated at compile time making them extremely fast.

* Hamming

* Isogram

* Leap

* List Ops

* Luhn                  - regexp mustcompile and matchstring, strconv.atoi converts a string to an int

* Pangram               - using a map to determine uniqueness as well as completeness (are all letters 
represented in a string)

* Proteins              - custom errors
  * type CustomError string
  * func (e CustomError) Error() string{ return string(e)}
  * const ( ErrInvalidBase = CustomError("error ... ") )
  * or just var ErrStop = errors.New("err stop")

* Proverb

* raindrops

* Reverse String

* RNA Transcription 

* Rotational Cipher       - if else/if else

* Scrabble Score

* Secret Handshake        - bit operations <<, >>, & (and), | (or), ^ (xor); use of uints for this purpose

* Spage Age

* Triangle

* Two Fer

* Word Count            - regexp mustcompile and findallstring; adding to an map without testing a key's 
existence first (for [string]int maps)