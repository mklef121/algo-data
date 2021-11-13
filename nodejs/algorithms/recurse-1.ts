function Factorial(num: number){

    if (num != 1) {
        return num * Factorial(num-1)
    }

    return num
}

console.log(Factorial(20));

