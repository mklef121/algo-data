import '../BigO/O(1).dart';

/// ##### Question: Given two arrays, create a function that lets a user compare if each array has a common Item.
///
/// Follow the rule that is on the cheat sheet
/// 1. When the interviewer says the question, write down the key points at the top (i.e. sorted array).
///  Make sure you have all the details. Show how organized you are.

List<String> arr1 = ["a", "x", "u", "o"];
List<String> arr2 = ["b", "q", "r", "p"];
// Should return false

List<String> arr3 = ["a", "x", "u", "o"];
List<String> arr4 = ["b", "q", "a", "p"];
//Should return true

/// 2. Make sure you double check: What are the inputs? What are the outputs?
/// The Inputs are arr1, arr2, arr3, arr4 and the expected output is boolean (`true` or `false`)

/// 3. What is the most important value of the problem? Do you have time, and space and memory,
///  etc.. What is the main goal?

//Since we are dealing with arrays that can grow quite large, It is very important that we treat time with caution
//Don't run any algorithm that will be O(n^2)

/// 4. Start with the naive/brute force approach. First thing that comes into mind. It shows that you’re able to think well and critically (you don't need to write this code, just speak about it)

bool arraCompare(List<dynamic> arrA, List<dynamic> arrB) {
  for (var count1 = 0; count1 < arrA.length; count1++) {
    for (var count2 = 0; count2 < arrB.length; count1++) {
      if (arrA[count1] == arrB[count2]) return true;
    }
  }

  return false;
}
//This is the brute force approach, explain why this would not work well

/// 5. Before you start coding, walk through your code and write down the steps you are going to follow.

//Write a new function
//turn the first array to a map Since look up in a map is constant time.
//Loop through the second array and compare it with the map

bool arrayCompare2(List<String> arrA, List<String> arrB) {
  var loadMap = Map<String, bool>();
  var newMap = arrA.fold(loadMap, (previousValue, element) {
    // print(previousValue);
    // print(element);
    var val = loadMap[element];
    // print("val is the $val");
    if (val == null) loadMap[element] = true;

    return loadMap;
  });

  for (var i = 0; i < arrB.length; i++) {
    if (newMap[arrB[i]] != null) return true;
  }
  // print(arrA);
  // print(loadMap);
  // print(newMap);
  return false;
}

main(List<String> args) {
  var vale = arrayCompare2(arr1, arr2);
  print(vale);
}
// 65k and 70k per annum
// 900 Euros a month rent
//
