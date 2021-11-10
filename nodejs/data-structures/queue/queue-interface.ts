import {LinkNode} from "../linked-list/linked-list"
export interface QueueInterface{
    peek()
    enqueue(value: any): this
    dequeue(value:any):LinkNode
}