
const nemo: string[]= ["nemo","Ultimate", "coding", "interview", "bootcamp", "Get", 
					   "more", "job", "offers", "negotiate", "a", "raise", "Everything", 
					   "you", "need", "to", "get", "the", "job", "you", "want"];

function loopArray(arr: Array<string>, index: number): string {
	return arr[index];
}

//This is an O(1), No matter how many elements are in the array
  // Just a single operation is performed
let res = loopArray(nemo,9);

console.log(res);

