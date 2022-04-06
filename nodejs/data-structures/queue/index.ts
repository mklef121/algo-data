import LinkedList, { LinkNode } from "../linked-list/linked-list";
import { QueueInterface } from "./queue-interface";

class Queue implements QueueInterface {
    first: LinkNode
    last: LinkNode
    length:number

    constructor(){
        this.first = null
        this.last = null
        this.length = 0
    }
    peek() {

       return this.first
    }
    enqueue(value: any): this {
        const node = new LinkNode(value)

        if (this.length === 0) {
            this.first = node
            this.last = node
        }else{
            node.setBehind(this.last)
            this.last.setNext(node)
            this.last = node
        }

        this.length ++
        return this
    }
    dequeue(): LinkNode {
        if (this.length === 0) {
            return
        }
        const toPop = this.first;

        if (this.length === 1) {
            this.first = null
            this.last = null

        }else{
            this.first = toPop.getBehind()
            toPop.setBehind(null)
        }
       
       this.length--
       return toPop
    }
    
}

const myQueue = new Queue()

myQueue.enqueue("Joy").enqueue("Miracle").enqueue("Titus")

console.log(myQueue,"\n\n\n", myQueue.peek());


myQueue.dequeue();
myQueue.dequeue();
myQueue.dequeue();
myQueue.dequeue();

console.log("\n\n\n\n",myQueue,"\n\n\n", myQueue.peek());

