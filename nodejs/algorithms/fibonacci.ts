
//`0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89,144, 233, 377,

function fibonacciIterative(n: number){
    const fibArra: number[] = [0,1];

    
    for (let index = 2; index <= n; index++) {
       
        fibArra.push(fibArra[index-1] + fibArra[index-2])
    }

    return fibArra[n]
}

//This has a Big O of O(2^n)
function fibonacciRecursive(n:number): number{
    if (n === 0) return 0;
    if (n === 1) return 1;

    return fibonacciRecursive(n-1) + fibonacciRecursive(n-2);
}

console.log(fibonacciRecursive(10),fibonacciIterative(10));