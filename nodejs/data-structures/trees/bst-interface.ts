import { TreeNode } from "./binary-search-tree";

export interface BstInterface{
    insert(data: any): this
    lookup(value: any):TreeNode
    remove(value:any): TreeNode
}