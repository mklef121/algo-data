List<int> arr = [3, 4, 5, 6, 7, 8];

int pickArrayValue(List<int> array, int index) {
  return array[index];
}

void main(List<String> args) {
  //This is an O(1), No matter how many elements are in the array
  // Just a single operation is performed
  var val = pickArrayValue(arr, 4);
  print(val);
}
