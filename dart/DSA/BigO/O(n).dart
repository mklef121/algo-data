List<String> names = [
  "nemo",
  "Ultimate",
  "coding",
  "interview",
  "bootcamp",
  "Get",
  "more",
  "job",
  "offers",
  "negotiate",
  "a",
  "raise",
  "Everything",
  "you",
  "need",
  "to",
  "get",
  "the",
  "job",
  "you",
  "want"
];
var arrayBig = List<String>.filled(900, "nemo");
main(List<String> args) {
  loopArray(names, "small array");
}

void loopArray(List<String> list, String description) {
  var beginNow = DateTime.now().millisecondsSinceEpoch;
  for (var i = 0; i < list.length; i++) {
    if (list[i] == "nemo") {
      // print("Found Nemo");
    }
  }
  var endNow = DateTime.now().millisecondsSinceEpoch;
  var milliSecDiff = endNow - beginNow;
  // print(DateTime.now().millisecondsSinceEpoch);
  print("The completion time for $description array is =>> $milliSecDiff");
}
