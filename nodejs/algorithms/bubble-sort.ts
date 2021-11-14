
const arr = [99,44,6,2,1,5,63,87,283,4,0]

function bubbleSort(arr: number[]): number[]{
    const length = arr.length;

    for (let index = 0; index < length; index++) {
        
        for (let secIndex = 0; secIndex < length; secIndex++) {
           if (arr[secIndex] > arr[secIndex+1]) {
               let first = arr[secIndex];
               arr[secIndex] =arr[secIndex+1];
               arr[secIndex+1] = first
           }
            
        }
        
    }

    return arr

}


console.log(bubbleSort(arr));
