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

* ISBN Verifier        - int (rune - '0') --- converts a rune number to the numeric value (ex '9' -> 9); strings.ReplaceAll(isbn, "-", "")

* Isogram

* Largest series product - 

* Leap

* List Ops

* Luhn                  - regexp mustcompile and matchstring, strconv.atoi converts a string to an int

* Nth Prime             - only have to go up to sqrt of a number to determine if it is prime

* Nucleotide Count      - if _, ok := myMap[v]; ok {} pattern to determine if 'v' is a valid
key for a map

* Pangram               - using a map to determine uniqueness as well as completeness (are all letters 
represented in a string)

* Parallel Letter Frequency - goroutines, channels; could have used an anonymous function instead of channelFrequency(); could have used buffered
channels (unnbuffered channel is usually used for signalling and synchronization). This will prevent goroutines from blocking when sending to the channel. However, the buffer doesn't need as many slots as there are goroutines, because we will be able to start receiving from the channel while some goroutines are still sending to it. So len(input) is not a good way to set the channel buffer. The exact optimal buffer size depends on the relative speed of the senders and receivers, and you can use trial and error to determine this, but for most purposes a small fixed size (say 10) is just fine.
Why do we need a channel buffer at all? Because if a goroutine wants to send its result but the main goroutine is not ready to receive it, the worker goroutine would block. To avoid that, we need enough channel slots that, if the rate of producing results exceeds the rate of consuming results, we can store the excess results until the consumer can catch up.  In practice, as long as the buffer is big enough to avoid too much blocking, it's fine. We wouldn't want to make it a function of N, because that would allow someone to crash the program by simply submitting a large enough N. The program would then try to allocate more memory than is available

* Phone Number          - use fmt.sprintf to format strings easily!  create the fundamental function first then create the others; regexp.replaceallstring

* Prime Factors         - using a for vs an if (see comments in go code for details)

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

* Tree Building        - sort.Slice(records, func(i, j int) bool {return records[i].ID < records[j].ID}), just need to provide Less function; using pointers in the children slice allows us to update each parent's children but only return the root node's pointer but it picks up the whole tree due
to the pointers in the children's children arays.

* Two Fer

* Word Count            - regexp mustcompile and findallstring; adding to an map without testing a key's 
existence first (for [string]int maps)