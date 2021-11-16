import { BstInterface } from "./bst-interface"

export class TreeNode {
    left?: TreeNode = null
    right?: TreeNode = null
    value: any
    constructor(val) {
        this.value = val
    }
}


class BinarySearchTree implements BstInterface {
    root: TreeNode = null
    constructor(parameters?: any) {
        if (parameters) {
            this.root = new TreeNode(parameters)
        }
    }
    remove(value: any): TreeNode {
        throw new Error("Method not implemented.")
    }
    insert(data: any): this {
        const node = new TreeNode(data)

        if (!this.root) {
            this.root = node
        } else {
            let rootNode = this.root

            while (true) {
                if (data < rootNode.value) {
                    //Left
                    if (!rootNode.left) {
                        rootNode.left = node
                        break
                    }

                    rootNode = rootNode.left
                } else {
                    //right
                    if (!rootNode.right) {
                        rootNode.right = node
                        break
                    }

                    rootNode = rootNode.right
                }
            }
        }
        return this
    }
    lookup(value: any): TreeNode {
        if (!this.root) return null

        let currentNode  = this.root;

        while (currentNode) {
            if (value < currentNode.value ) {
                currentNode = currentNode.left
            }else if (value > currentNode.value ) {
                currentNode = currentNode.right
            }else{
                break
            }
        }
       return currentNode
    }

    // remove

    breathFirstSearch(value: any): TreeNode | null {
        let queue: TreeNode[] = [this.root]
        let result = []
        while (queue.length > 0) {

            let currentNode = queue.shift();
            // console.log(currentNode.value);
            result.push(currentNode)
            // if (currentNode.value == value) {
            //     queue = []
            //     return currentNode
            // }

            if (currentNode.left) {
                queue.push(currentNode.left)
            }

            if (currentNode.right) {
               queue.push(currentNode.right) 
            }

        }
        return null
        
    }
//             56
//       45          87
//     24  48     70   98
//   23  25
    BFSrecursion(){
        // We are expected to get
        // [56,45,87,24,48,70,98,23,25]
        let queue = [this.root]
        
        function recurseBFS(result: any[]){

            if (!queue.length) {

                return result
            }
            let current = queue.shift();
           
            result.push(current.value);
            
            if (current.left) {
                queue.push(current.left)
            }

            if (current.right) {
                queue.push(current.right)
            }

            return recurseBFS(result)
        }

        return recurseBFS([])
    }

    depthFirstSearchInOrder(){
        // [56,45,24,23,25,48,87,70,98]

        return inOrderTransverse(this.root, [])
    }

    depthFirstSearchPreOrder(){
        //Expected 
        // [56,45,24,23,25,48,87,70,98]
        let result = []
        return preOrderTransverse(this.root, result)
        // return result 
    }

    depthFirstSearchInPostOrder(){

        return postOrderTransverse(this.root, [])
    }
}

function preOrderTransverse(current: TreeNode, TransVersedlist: any[]): any[]{

    TransVersedlist.push(current.value)
    if (current.left) { 
        inOrderTransverse(current.left, TransVersedlist)
    }
    
    if (current.right) {
        inOrderTransverse(current.right, TransVersedlist)
    }

    return TransVersedlist
   
}


function inOrderTransverse(current: TreeNode, TransVersedlist: any[]): any[]{
    //Idea is to have something like this
    // [23,24,25,45,48,56,70,87,98]
     if (current.left) {
        inOrderTransverse(current.left, TransVersedlist)
     }

     TransVersedlist.push(current.value)

     if (current.right) {
        inOrderTransverse(current.right, TransVersedlist)
     }

     return TransVersedlist
}

function postOrderTransverse(current: TreeNode, TransVersedlist: any[]): any[]{
    // The idea is to have a final looking thus
    // [23,25,24,48,45,70,98,87,56]
    if (current.left) {
        postOrderTransverse(current.left, TransVersedlist)
    }

    if (current.right) {
        postOrderTransverse(current.right,TransVersedlist)
    }

    TransVersedlist.push(current.value)

    return TransVersedlist
}

//Our BST looks like this

//             56
//       45          87
//     24  48     70   98
//   23  25

function traverse(node) {
    const tree = { value: node.value, left:null, right:null};
    tree.left = node.left === null ? null : traverse(node.left);
    tree.right = node.right === null ? null : traverse(node.right);
    return tree;
  }


const bst = new BinarySearchTree()
bst.insert(56);
bst.insert(45);
bst.insert(87)
.insert(70)
.insert(24)
.insert(48)
.insert(98)
.insert(23)
.insert(25)
console.log(JSON.stringify(traverse(bst.root)));

const found =  bst.lookup(984)

console.log("\n\n\n",found);

console.log("\n\n","BFS result \n", bst.breathFirstSearch(70));
 
console.log("\n\n",bst.BFSrecursion());

//Our BST looks like this

//             56
//       45          87
//     24  48     70   98
//   23  25

