let undortedArray = [47,3,4,10,5,6,7,7,1,9,23,8,2];

//This is a horrible algorithm
// I wish to re-write it later
function bubbleSort(arr: number[]){

    let reRange = true;

    while (reRange) {
        console.log("We count");
        
        reRange= false;
        for (let index = 1; index < arr.length; index++) {
            const element = arr[index];
            const formerIndex = index-1;

            if (arr[formerIndex] >  arr[index]) {
                let former = arr[formerIndex];
                let current = arr[index];
                arr[formerIndex] = current;
                arr[index] = former;
                reRange = true;

            }
            
        }
    }

    return arr;
}


console.log(bubbleSort(undortedArray))