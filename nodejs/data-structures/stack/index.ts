import {LinkNode} from "../linked-list/linked-list"
import { StackInterface } from "./stack-interface"

export class Stack implements StackInterface{
    public length: number
    top ?: LinkNode
    bottom ?: LinkNode
    constructor(){
        this.top = null
        this.bottom = null
        this.length = 0
    }

    peek(){
        return this.top
    }

    push(value: any){
        const node = new LinkNode(value)
        if(this.length == 0){
            this.top = node;
            this.bottom = node;
        }else{
            // this.top.setNext(node);
            node.setBehind(this.top)
            
            // node.setNext(this.top);
            this.top = node;
        }

        this.length ++;

        return this;
    }

    pop(){
        if (this.top) {
           const tempNode =  this.top;
           this.top = tempNode.getBehind()
           this.length --

            if (this.length === 0) {
                this.bottom = null
            }
        }

        return this
    }

    isEmpty(){
        return this.length === 0
    }
}

let mystack = new Stack();

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
