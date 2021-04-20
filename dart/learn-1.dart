//Vartiable
var name = 'Bob';
Object name1 = 'Bob';
String name2 = 'Bob';

// Default value
// Dart uses null safety fromversion 2.2, so for a value to have a default value of NULL, declare it thus
int? lineCount;
// assert(lineCount == null); -- returns true

//this is used to initialize a variable that cannot be nullable
late String description;

//Final and const
//If you never intend to change a variable, use final or const,
final finalName = 'Bob'; // Without a type annotation
final String nickname = 'Bobby';

//Use const for variables that you want to be compile-time constants. Thus cannot be instance type(or class) members
const baz = [3, 4, 5]; // Equivalent to `const []`

// You can define constants that use type checks and casts (is and as), collection if, and spread operators (... and ...?):
// I'll read and understand these better
const Object i = 3; // Where i is a const Object with an int value...
const list = [i as int]; // Use a typecast.
const map = {if (i is int) i: "int"}; // Use is and collection if.
const set = {if (list is List<int>) ...list}; // ...and a spread

void main() {
  // int two;
  description = 'Feijoada!';
  int lineCount2;
  if (true) {
    lineCount2 = 3;
  } else {
    lineCount2 = 0;
  }

  print(lineCount2);
  print(lineCount);
  print(description);
  print(baz);
  print(map);
}
