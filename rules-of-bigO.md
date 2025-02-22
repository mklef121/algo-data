### What is Big O?

Big O notation is a system for measuring the rate of growth of an algorithm. Big O notation mathematically describes the complexity of an algorithm in terms of time and space.

The O is short for “Order of”. So, if we’re discussing an algorithm with O(log N), we say its order of, or rate of growth, is “log n”, or logarithmic complexity.

### O(log n) 

Logarithms allow us to reverse engineer a power. (Like Kryptonite!) They are the inverse operation of exponentiation.
We can think of logarithms as the opposite operation of exponentiation.

remember 
since 1000 = 10 × 10 × 10 = 103
but 
log10 (1000) = 3
The logarithm of x to base b is denoted as logb (x), or without parentheses, logb x


#### Rules to define the Complexity of BigO

There are several rules to find how the BigO of a function or a code since several scenarios might be playing out
in the said function.


- Rule 1: **The Worst Case**
Assume that the function operates at it's worst scenario. E.G

```go
package main;

hold:= []string{
    "hi","come","Nemo","Dull","Rahman","Django","Tall","Build"
}

func main(){
    for  i:=0; i < len(hold); i++ {
        if hold[i] == "Nemo"{
            break;
        }
    }
}

//if this main function runs, we can say it's Big O(3), but worst case scenario will assume the array is larger
// Than what it is currently, and that `Nemo` is the last element of the array
```
Thus we will assume that this is Big O(n)

- Rule 2: **Remove Constants**

A function might conntain several constant operations also known as O(1). if other complex operations are used, then remove the constants.
```dart
void main(){
   List<String> arr = ["A", "function", "might", "contain", "several", "constant", "operations", "also", "known", "as"];

   manipulate(arr);
}

void manipulate(List<String> arg){

    var first = arg[0]; //O(1)

    var newOne = "First $first"; //O(1)

    //print all names
    arg.map((title) => print(title) ); //O(n)

    arg.map((title) => print("The title is duplicated as $title") ); //O(n)
    
    int count = 200;

   for(var i = 0; i < count; i++){ //O(200)
        print("Constant loop of ${arg[i]}");
    }
}
```
Normally, the function above would have been O(n + n + 202) == O(2n + 202) but to abide by the rule, we remove the constant operations
--> It becomes O(n);

- Rule 3: **Different terms for input**

We might have a function taking multiple parameters, the big O should take into cognisance that the effect of the two inputs

```ts
let arr1 = ["A", "function", "might", "contain", "several", "constant", "operations", "also", "known", "as"];
let arr2 = ["constant", "operations", "also", "known", "as", "A", "function", "might", "contain", "several", ];

function manipull(arr1, arr2){
    arr1.ForEach((item) => console.log(item)) //O(n)

    arr2.ForEach((item) => console.log(item)) //O(m)
}

manipull(arr1, arr2);
```
The Big O will be O( m + n); this is because of the rule. We must take inputs differently because they are looping through separate Items.

- Rule 4: **Drop Non Dominants**

In a function, there may be two or more occurences of several Big O types. But its good to note the level at which
the complexities of Big O increases which are 
 * O(1) Constant- no loops
 * O(log N) Logarithmic- usually searching algorithms have log n if they are sorted (Binary Search) 
 * O(n) Linear- for loops, while loops through n items
 * O(n log(n)) Log Liniear- usually sorting operations
 * O(n^2) Quadratic- every element in a collection needs to be compared to ever other element. Two nested loops
 * O(2^n) Exponential- recursive algorithms that solves a problem of size N
 * O(n!) Factorial- you are adding a loop for every element

When several of them are combined in a function, drop the ones at the top of the list(They are less dominant)

```go
func printNumbersAndPairSums(arr []int){
    for _,num := range arr{
        fmt.Println(num)  //O(n)
    }

    //prints a sum of each number against all numbers in the array
    for _,num := range arr{
        for _,num2 := range arr{
         fmt.Println(num + num2) //O(n^2)
        }
    }
}

```
Looking at this now, we will think its Big O will be BigO(n + n^2), but it's actuallly **BigO(n^2)** because it we dropped the non dominant notation

##### Calculating O(log n): Binary Search

The classic example used to illustrate O(log n) is binary search. Binary search is an algorithm that finds the location of an argument in a sorted series by dividing the input in half with each iteration.

EXAMPLE

Let’s say we are given the following array and asked to find the position of the number `512`:

```ts
const powers = [1, 2, 4, 8 ,16, 32, 64, 128, 256, 512];
```

First, let’s review the brute force solution to this problem.

```ts
const bruteSearch = (arr, num) => {
  
   for (let i = 0; i < arr.length; i++) {
       if (arr[i] === num) {
           return `Found ${num} at ${i}`;
       }
   }
}
```

What’s the Big O of bruteSearch()?

`O(n)`

Can we do better?

Definitely, Lets Use Binary Search

```ts
const powers = [1, 2, 4, 8 ,16, 32, 64, 128, 256, 512];
// find Position of 512
const binarySearch = (arr, num) => {
 
   let startIndex = 0;
   let endIndex = (arr.length)-1;
//    0, 9
// 4.5 //4
// 5, 9 // 14/2 = 7
// 8, 9 // 17/2 
  
   while (startIndex <= endIndex){
      
       let pivot = Math.floor((startIndex + endIndex)/2);
 
       if (arr[pivot] === num){
            return `Found ${num} at ${pivot}`;
       } else if (arr[pivot] < num){
           startIndex = pivot + 1;
       } else {
           endIndex = pivot - 1;
       }
   }
   return false;
}
```


