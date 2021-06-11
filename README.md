Golang exercises through exercism.io

What I've learned from each exercise

* Accumulate            - ability to pass functions as a parameter to a function

* All your base         - outputDigits = append([]int{divisionResult%outputBase}, outputDigits... )

* Anagram               - strings.EqualFold for case insensitive string equality and reflect.DeepEqual for map equality

* Atbash Cipher         - interacting with runes like their ASCII integers (use them in math equations, etc.)

* Bob                   - regexp mustcompile and matchstring, switch/case

* Clock                 - pointer vs value receivers (https://github.com/golang/go/wiki/CodeReviewComments#receiver-type)

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

* Nucleotide Count      - if _, ok := myMap[v]; ok {} pattern to determine if 'v' is a valid
key for a map

* Pangram               - using a map to determine uniqueness as well as completeness (are all letters 
represented in a string)

* Phone Number          - use fmt.sprintf to format strings easily!  create the fundamental function first then create the others; regexp.replaceallstring

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

* Sieve                   - make([]bool, 5) makes a list of 5 'falses'; go doesn't have a func to test slice membership

* Simple Cipher           - using structs to implement a generic interface where there is shared behavior but differences in details.  Using the patter of a "New()" function to return a new struct
generic interface allows code reuse since you can refer to different struct implementations as the
inherited interface.

* Spage Age

* Strain               - passing predicate functions as arguments to another function, to perform actions on elements of a slice

* Triangle

* Two Fer

* Word Count            - regexp mustcompile and findallstring; adding to an map without testing a key's 
existence first (for [string]int maps)