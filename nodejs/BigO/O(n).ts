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

const nemo: string[]= ["nemo","Ultimate", "coding", "interview", "bootcamp", "Get", 
					   "more", "job", "offers", "negotiate", "a", "raise", "Everything", 
					   "you", "need", "to", "get", "the", "job", "you", "want"];
const arr_use: Array<string> = new Array(30000).fill("nemo");

function findNemo():void{

	loopArray(nemo, "small")
	// loopArray(arr_use, "Large")
    
}

/**
 * This function runs with a linear BigO notation **O(n)**
 * This means as the size of the array increases, then the number of operations increases
 * This is Linear time
 * @param arr {Array}
 * @param description 
 */
function loopArray(arr, description: string) {
	let t0:number  = performance.now();
    for (var i = 0; i < arr.length; ++i) {
    	if (arr[i] == 'nemo') {
			// one notable factor was how long the console command was delaying the processing time
    		// console.log("Found nemo at index: ", i)
    	}
    }

    let t1:number  = performance.now();

    console.log("The call to find Nimo took %s milli seconds for the %s array", t1-t0, description)
}

findNemo();
