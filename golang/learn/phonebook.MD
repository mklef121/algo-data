## PhoneBook Application 

This utility offers support for two commands named search and list that search for a given surname and return its full record if the surname is found, and lists all available records, respectively.

The Steps to implement the version 1 of the phonebook app will be thus


- There exists a new user-defined data type for holding the records of the phone book that is a Go structure with three fields named Name, Surname, and Tel. Structures group a set of values into a single data type, which allows you to pass and receive this set of values as a single entity.
- There exists a global variable that holds the data of the phone book, which is a slice of structures named data.
- There exist two functions that help you implement the functionality of the search and list commands.
- The contents of the data global variable are defined in the main() function using multiple append() calls. You can change, add, or delete the contents of the data slice according to your needs.
- Lastly, the program can only serve one task at a time. This means that to perform multiple queries, you have to run the program multiple times.