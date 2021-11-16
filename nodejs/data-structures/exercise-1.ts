/**
 * You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

 

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */
 class ListNode {
     val: number
     next: ListNode | null
     constructor(val?: number, next?: ListNode | null) {
         this.val = (val===undefined ? 0 : val)
         this.next = (next===undefined ? null : next)
     }
 }


 function mergeKLists(lists: Array<Array<number> | null>): Array<number> | null {
    if(lists.length == 0){
        return null
    }

    let allListArray: number[];
    for(let i = 0; i < lists.length; i++){
        if(lists[i]){
            // let member:number[] = getArrayFromLink(lists[i])
            if(allListArray) {
                allListArray = mergeAndSortTwoArrays(allListArray, lists[i])
            }else{
                allListArray = lists[i]
            }
           
        }
    }
    
    return allListArray;
};

function turnArrayToLinkedNode(array: number[]): ListNode{
    let node =  new ListNode(array[0])
    let lastNode = node
    for(let i = 1; i< array.length; i++){
        let newNode = new ListNode(array[i])
        lastNode.next = newNode
        lastNode = newNode
    }

    return node
}


function mergeAndSortTwoArrays(left: number[], right: number[]): number[]{
    
    let newList:number[] = []
    
    let leftIter = 0
    let rightIter = 0
    let currentLeft = left[leftIter]
    let currentRight = right[rightIter]
    

    while(leftIter < left.length || rightIter < right.length){
        if(currentLeft < currentRight && leftIter < left.length || rightIter >= right.length){
          
            newList.push(currentLeft);
            leftIter ++
            if(leftIter < left.length)
             currentLeft = left[leftIter];
            
        }else{
            newList.push(currentRight)
            rightIter++
            
            if(rightIter < right.length)
             currentRight = right[rightIter]
        }
    }

    return newList
}

function getArrayFromLink(linked: ListNode):number[] {
    let next = linked.next;
    let array:number[] = []
    while(next){
        array.push(next.val)
        next = linked.next;
    }
    return array;
}


console.log(mergeKLists([[1,4,5],[1,3,4],[2,6]]));
