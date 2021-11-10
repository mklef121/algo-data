import {LinkNode} from '../linked-list/linked-list'
export interface StackInterface{
    peek(): LinkNode
    push(value: any):this


pop(value: any): this
isEmpty(): boolean
}