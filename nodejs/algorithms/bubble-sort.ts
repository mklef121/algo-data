
const arr = [99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0]

function bubbleSort(arr: number[]): number[] {
    const length = arr.length;

    for (let index = 0; index < length; index++) {

        for (let secIndex = 0; secIndex < length; secIndex++) {
            if (arr[secIndex] > arr[secIndex + 1]) {
                let first = arr[secIndex];
                arr[secIndex] = arr[secIndex + 1];
                arr[secIndex + 1] = first
            }

        }

    }

    return arr

}

function secondBubbleSort(arr: number[]): number[] {

    for (let i = 0; i < arr.length; i++) {
        for (let j = 0; j < arr.length; j++) {
            if (arr[j] > arr[i]) {
                let first = arr[i];
                arr[i] = arr[j]
                arr[j] = first
                continue
            }
        }
    }

    return arr;
}


console.log("Before Sorting", arr);

console.log("\n\n After sorting", secondBubbleSort(arr));
