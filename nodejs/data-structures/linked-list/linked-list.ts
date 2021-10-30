
export default class LinkedList{
    //The first Item in the linked list and a pointer to it's descendedants
    head?: LinkNode;

    //The last item in the Linked list
    tail?: LinkNode;
    length: number
    constructor(value){
      this.head = new LinkNode(value);
      this.tail = this.head;
      this.length  = 1;
    }

    append(value: any){
        let newNode = new LinkNode(value)
        // On the first iteration, this is equivalent to this.head.next, thus creates the contigous chain
        this.tail.setNext(newNode);
        this.tail = newNode;

        this.length++

        return this
    }

    prepend(value: any){
        let newNode = new LinkNode(value)
        newNode.setNext(this.head);

        this.head = newNode;
        this.length ++;
        return this;
    }

    insert(index: number, value: any){
        if (index === 0) {
            return this.prepend(value);
        }

        if (index >= this.length) {
            return this.append(value)
        }
        let nextNode = this.head.getNext();
        let previous = this.head;

        let i = 1
        while (nextNode) {
            if (i === index) {
                let currentNode =  new LinkNode(value);

                currentNode.setNext(nextNode)
                previous.setNext(currentNode)
                this.length ++
                break
            }

            i++
            previous = nextNode
            nextNode = nextNode.getNext();
        }
        
    }

    remove(index: number){
        if (index === 0 && this.length === 1) {
            this.head = undefined
            this.tail = undefined
            return this
        }
        
        let nodeToDelete = this.head.getNext();
        let previous = this.head;

        let i = 1
        while (nodeToDelete) {
            if (i === index) {
               previous.setNext(nodeToDelete.getNext());

               //if it's last element
               if (i === (this.length -1)) {
                   this.tail = previous;
               }
               this.length --
               break
            }

            i++
            previous = nodeToDelete
            nodeToDelete = nodeToDelete.getNext();
        }
        

        return this;
       
    }

    reverse(){
        
    }

    printList(){
        const arr : any[] = []
        let current = this.head;
        while (current) {
            arr.push(current.getValue())

            current = current.getNext()
        }

        return arr
    }
}


class LinkNode{
   private next?:LinkNode 

    constructor(private value: any){
    }

    setNext(nxt:LinkNode ){
        this.next = nxt;
    }

    getValue(){
        return this.value
    }

    getNext(){
        return this.next
    }
}

let link =  new LinkedList("7748");
link.append("4")
link.append(5)
link.append(56).prepend("566")

link.insert(3,"nnenna")
link.insert(5,"Mmesoma")

console.log("Before remove",link.printList())
link.remove(3)
console.log("After remove Node 3",link.printList())