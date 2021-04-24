// export {}
import { performance, PerformanceObserver } from 'perf_hooks';


// What is good code ?

/*
Good code can be described in terms of 

1. Readability : can others understand your code ?
2. Scalabilty (Big O helps us measure this)

BigO notation is the language we use to talk about how long our algorith takes to run

BigO measures how many operations the computer will perform in running our code if we 
increase the number of elements involved 
*/

const nemo: number[]= [1,2,3,4,5,6,7,14,19,23];


/**
 * This function runs with a linear BigO notation **O(n * n) =>> O(n^2)**
 * This means as the size of the array increases, then the number of operations increases
 * This is Quadratic time
 * @param arr {Array}
 * @param description 
 */
function loopArray(arr: number[]) {
	let t0:number  = performance.now();
    for (var i = 0; i < arr.length; ++i) {
    	for (let index = 0; index < arr.length; index++) {
			console.log(arr[i]+" "+arr[index]);
		}
    }

    let t1:number  = performance.now();

    console.log("The call to loop this array took took %s milli seconds", t1-t0)
}

loopArray(nemo);
