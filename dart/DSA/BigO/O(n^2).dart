List<int> names = [1, 2, 3, 4, 7, 8, 4, 10, 13, 65];
main(List<String> args) {
  formPair(names);

  // The function above runs two loops which we can call a multiplication
  // O(n * n) =>> O(n^2)
}

void formPair(List<int> list) {
  var beginNow = DateTime.now().millisecondsSinceEpoch;
  int count = list.length;
  for (var i = 0; i < count; i++) {
    for (var j = 0; j < count; j++) {
      print("${list[i]} ${list[j]} ");
    }
  }
  var endNow = DateTime.now().millisecondsSinceEpoch;
  var milliSecDiff = endNow - beginNow;

  print(
      "The completion time  for the looping array is =>> $milliSecDiff millisecond");
}
