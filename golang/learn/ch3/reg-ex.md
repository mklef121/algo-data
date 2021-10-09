# Regular expressions and pattern matching

**Pattern matching** is a technique for searching a string for some set of characters based on a specific search pattern that is based on regular expressions and grammars.

A **regular expression** is a sequence of characters that defines a search pattern

> Every regular expression is compiled into a recognizer by building a generalized **transition diagram called a finite automaton**.

A recognizer is a program that takes a string x as input and is able to tell whether x is a sentence of a given language.

A finite automaton can be either deterministic
or nondeterministic. Nondeterministic means that more than one transition out of a state can be possible for the same input.

####  Metacharacters
A metacharacter is a character that has a special meaning during pattern processing. You use metacharacters in regular expressions to define the search criteria and any text manipulations.


1.	`ˆ`  : lets the recognition start at the beginning or negation when used in the beginning of [ ] -character class 
2.	`$`  : lets you define where the pattern ends.
3.	`*`   : is a placeholder that means no or any number of characters.
4.	`+`  : is a placeholder that means one or any number of characters.
5.	`?`  : is a placeholder that means no or one character.
6.	`[a-z]`  : defines one character out of group of letters or digits. Only a single character from the group is used
7.	`()`   : groups characters or strings of characters. You can use the set operators *, +, and ? in such a group, too. It means all the characters generated from the group is one.
8.	`{}`  : Is a repetition marker that defines the character before the braces. The range can be defined by numbers; if the start and the end are given separately, the numbers are written with a comma ({3,7}) 
9.	`\` (the backslash) : Masks metacharacters and special characters so that they do no longer possess a special meaning.
10.	`.` : Represents exactly any character. If you actually need a dot, just write `\.`
11.  `\t`  -------  Tabulator
12.  `\n` -------   New line
13.  `\r` --------  Return (carriage return)
14.  `\f` --------  Form feed
15.  `\v` --------  Vertical tabulator
16.  `\s` --------  White space (not visible in print, includes \t, space, \n, \r, \f)
17.  `\S` ---------  Negation of \s
18.  `\w` -------    Word character (letters that form words, such as in [_a-zA-Z0-9]) 
19.  `W` -------    The negation of \w
20.  `\d` --------- Digits, like [0-9]
21.  `\D` ---------  Negation of \d
22.  `\b` --------- Word boundary, start, or end of a word; all not in the \w definition. 
23.  `\B` ---------   Negation of \b
24.  `\0` --------    Nul (nil) character (physical 0)
25.  `\xxx` ------    Character value, written as an octal number
26.  `\xdd` ------   Character value, written as an hexadecimal number
27.  `\uxxxx` ----   Unicode character, written as hex number
28.  `\cxxx`  -----  Control, ASCII value