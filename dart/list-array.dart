List<int> list = [1, 2, 3]; //returns a List Class
// assert(list.length == 3);
// assert(list[1] == 2);

// To create a list that’s a compile-time constant, add const before the list literal:

var constantList = const [1, 2, 3];
// constantList[1] = 1; // This line will cause an error.
//
// using spreat operator
var list2 = [0, ...list];
// assert(list2.length == 4);

List<int>? list3;

//  the expression to the right of the spread operator might be null, you can avoid
// exceptions by using a null-aware spread operator
var list4 = [...list2, ...?list3];
bool promoActive = false;
var nav = ['Home', 'Furniture', 'Plants', if (promoActive) 'Outlet'];

var listOfInts = [1, 2, 3];
var listOfStrings = ['#0', for (var i in listOfInts) '#$i'];
// assert(listOfStrings[1] == '#1');

main(List<String> args) {
  var list34 = list;
  list34[0] = 56;
  print("List One is $list");
  print("List 34 copied is $list34");
  print(list2);
  print(list4);
  print(nav);
}

//generics
abstract class Cache<T> {
  T getByKey(String key);
  void setByKey(String key, T value);
}

enum Size { small, medium, large }

T stringToEnum<T>(String str, Iterable<T> values) {
  return values.firstWhere(
    (value) => value.toString().split('.')[1] == str,
    // orElse: () => null,
  );
}

Size size = stringToEnum<Size>("medium", Size.values);
