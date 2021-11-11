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
        throw new Error("Method not implemented.")
    }

    // remove
}

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
