

// What is good code ?

/*
Good code can be described in terms of 

1. Readability
2. Scalabilty (Big O helps us measure this)
*/

const nemo: string[]= ["nemo","Ultimate", "coding", "interview", "bootcamp", "Get", 
					   "more", "job", "offers", "negotiate", "a", "raise", "Everything", 
					   "you", "need", "to", "get", "the", "job", "you", "want"];
const arr_use: Array<string> = new Array(2000).fill("Mewo");

function findNemo(arr):void{

	let t0:number  = performance.now();
    for (var i = 0; i < arr.length; ++i) {
    	if (arr[i] == 'nemo') {
    		console.log("Found nemo", i)
    	}
    }

    let t1:number  = performance.now();

    console.log("The call to find Nimo took %s milli seconds", t1-t0)
    
}

findNemo(nemo);
