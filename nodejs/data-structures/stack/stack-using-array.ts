import { StackInterface } from "./stack-interface"

class ArrayStack implements StackInterface{
    data:any[] = []
    push(value: any): this {
        this.data.push(value)
      return this
    }
    pop(): this {
        this.data.pop()

        return this
    }
    

    peek(){
        return this.data[this.data.length -1]
    }

    isEmpty(){
        return this.data.length === 0
    }
}

let mystack = new ArrayStack();

console.log(mystack.peek(), "The Initial");

mystack.push("google")

console.log(mystack.peek(), "\n\n The Subsequesnt", mystack);

mystack.push("Udemy")

console.log(mystack.peek(), "\n\nThe Next", mystack);

mystack.push("Cosera")

console.log(mystack.peek(), "\n \nThe Next Course \n", mystack);

mystack.pop();
mystack.pop();
mystack.pop();

console.log("\n\nThe main stack \n", mystack);