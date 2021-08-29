
// "Miracle" => elcariM
function stringReverse(str :string): string{

    return str.split("").reverse().join("");
}

console.log(stringReverse("Miracle"));


//The idead [1,2,3,4,6,8] and [4,5,7,9] will produce [1,2,3,4,5,6,7,8,9]
function mergeSortedArrays(arr1: number[],arr2: number[]): number[]{
    let mergedArray: number [] = [];

    let first = arr1[0];
    let second = arr2[0];
    // const newArray = [];
    let totalIteration = 0;
    let i = 0;
    let j = 0;
    let toCompeteB = 0;

    while (first != undefined || second != undefined) {
        toCompeteB = second;
        if (second == undefined) toCompeteB = Infinity;

        if ((first != undefined) && first < toCompeteB) {
            mergedArray.push(first)
            i++;
            first = arr1[i];
            // console.log(first,"first did it", i)
        }else{
            mergedArray.push(second);
            j++;
            second = arr2[j];
            // console.log(second,"second did it", j)
        }
        
        totalIteration++
    }

    console.log("The entire iteration is", totalIteration)

    return mergedArray;
}


let resp = mergeSortedArrays([1, 2, 3, 7, 9, 23, 78],[5, 6, 8, 45, 67, 70]);

//gives O(a+b)

console.log(resp);
