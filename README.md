# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.


Use go build compoundword.go or go run compoundword.go.

The program will look for a file in the same directory named word.list by default.
Use -f filename to specify a different file.
-t will cause the program to use the trie based implementation instead of the original.
-time will cause the program to also output the run time of the implementation but does
not include the time to load the list from the file.
