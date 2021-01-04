

// Boolean

let isDone: boolean = false;


//Number

let decimal: number = 6;
let hex: number = 0xf00d;
let binary: number = 0b1010;
let octal: number = 0o744;
let big: bigint = 100n;


// String
let color: string = "blue";
color = 'red';


//Array

var list: number[] = [1, 2, 3];
var list: Array<number> = [1, 2, 3];


//Tuple

// Tuple types allow you to express an array with a fixed number of 
// elements whose types are known, but need not be the same.

// Declare a tuple type
let erx: [string, number];
// Initialize it
erx = ["hello", 10]; // OK
// erx= [200n,"fish",500]


// Enum

// an enum is a way of giving more friendly names to sets of numeric values


enum Color {
  Red,
  Green,
  Blue,
}
let c: Color = Color.Green;

console.log(c) //Green


// By default, enums begin numbering their members starting at 0

enum AnotherColor {
  Red = 2,
  Green,
  Blue,
}

var blu:string = AnotherColor[3]; //green


// Unknown

// We may need to describe the type of variables that we do not know when we are writing an application. 
// These values may come from dynamic content 

let notSure: unknown = 4;
notSure = "maybe a string instead";

// OK, definitely a boolean
notSure = false;



// ANY

// we might want to opt-out of type checking. To do so, we label these values with the any type:

declare function getValue(key: string): any; //A package wrriten by someone.
// OK, return value of 'getValue' is not checked
const str: string = getValue("myString");


// Void

// the absence of having any type at all

function warnUser(): void {
  console.log("This is my warning message");
}



// Null and Undefined

let u: undefined = undefined;
let n: null = null;

// By default null and undefined are subtypes of all other types.


// Never

// The never type represents the type of values that never occur.


// Function returning never must not have a reachable end point
function error(message: string): never {
  throw new Error(message);
}

// Inferred return type is never
function fail() {
  return error("Something failed");
}

// Function returning never must not have a reachable end point
function infiniteLoop(): never {
  while (true) {}
}



// Object

// This represents non primitive types. i.e: anything that is not number, string, boolean, bigint, symbol, null, or undefined.

declare function create(o: object): void;

// OK
create({ prop: 0 });

create(null);


// Type assertions

// Sometimes you’ll end up in a situation where you’ll know more about a value than TypeScript does.



let someValue: unknown = "this is a string";

let strLength: number = (someValue as string).length;

// or

let striLength: number = (<string>someValue).length;










