//array operations
// import '../../learn-1.dart';

List<String> test = ["she", "he", "her"];

main(List<String> args) {
  unshift<String>(test, "hero");
  slice(test, 4, 5);
  print(test);
}

int unshift<T>(List<T> list, [element]) {
  if (element is List<T>) {
    list.insertAll(0, element);
  } else if (element is T) {
    list.insert(0, element);
  }

  return list.length;
}

T pop<T>(List<T> list) {
  return list.removeAt(list.length);
  //or return list.removeLast();
}

T shift<T>(List<T> list) {
  return list.removeAt(0);
}

List<T> slice<T>(List<T> list, int start, [int? end]) {
  end = end == null ? list.length : end;
  return list.sublist(start, end);
}

//unshift()
//shift();
//slice()
//pop()
//push()
