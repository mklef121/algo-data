

//Implementing Fibonachi using dynamic programming
// 0,1,1,2,3,5,8,13,21,34,55,89 e.t.c

// before

function FibNumber(n: number){
    if (n == 0) return 0
    if (n <= 2) return 1

    return FibNumber(n-1) + FibNumber(n-2);
}

console.log(FibNumber(8));

//Implementing Fibonachi with Caching
function dynamicFibonacci(){

    let cache = {}

    return function fibNumber(n){
        if(n in cache) return cache[n]

        if (n<2) {
            return n
        }

        cache[n] = fibNumber(n-1) + fibNumber(n-2)

        return cache[n]
    }
}

const memoFib = dynamicFibonacci()

console.log(memoFib(60));

