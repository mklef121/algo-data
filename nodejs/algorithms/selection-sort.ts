const sortArray = [99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0]
function selectionSort(arr: number[]): number[] {

    let length = arr.length;

    for (let index = 0; index < length; index++) {
        let min = index;
        let tempVal = arr[index]

        for (let j = index+1; j < length; j++) {
            if (arr[j] < arr[min]) {
                min = j
            }
        }

            arr[index] = arr[min]
            arr[min] = tempVal

    }

    return arr;

}

function selectionSort2(array) {
    const length = array.length;
    for(let i = 0; i < length; i++){
      // set current index as minimum
      let min = i;
      let temp = array[i];
      for(let j = i+1; j < length; j++){
        if (array[j] < array[min]){
          //update minimum if current is lower that what we had previously
          min = j;
        }
      }
      array[i] = array[min];
      array[min] = temp;
    }
    return array;
  }
  

console.log(selectionSort(sortArray));
