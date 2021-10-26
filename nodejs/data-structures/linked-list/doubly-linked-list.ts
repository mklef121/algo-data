class DoublyLinkedList{
     //The first Item in the linked list and a pointer to it's descendedants
     head?: LinkNode;

     //The last item in the Linked list
     tail?: LinkNode;
     length: number
    constructor(value: any){
        this.head = new LinkNode(value);
        this.tail = this.head;
        this.length = 1;
    }

    append(value: any){
       let node = new LinkNode(value);
       node.setPrevious(this.tail);
       this.tail.setNext(node);
       this.tail  =  node;
        this.length++;
       return this;
    }

    prepend(value: any){
        let node = new LinkNode(value);
        node.setNext(this.head);
        this.head?.setPrevious(node);

        this.head = node;
        this.length ++
        return this;
    }

    insert(index: number, value: any){
        let found = this.findNode(index);
        let newNode = new LinkNode(value); 

        if (found) {
            
            newNode.setNext(found);
            newNode.setPrevious(found.getPrevious());
            found.getPrevious()?.setNext(newNode);
            found.setPrevious(newNode)
            this.length ++;
        }
        
       return this;
    }

    remove(index: number){
        let found = this.findNode(index);

        if (found) {
            found.getNext()?.setPrevious(found.getPrevious())

            found.getPrevious()?.setNext(found.getNext())

            this.length --
        }

        return this
    }
    protected findNode(index: number){
        let lastItemIndex = this.length -1;
        if (index > lastItemIndex || index <0) return
        
        let halfLength = lastItemIndex/2;
        let current: LinkNode = this.head

        //indicates if finding a node starts from the head
       let startFromHead: boolean = true
       let count = 0;
        if (index > halfLength) {
            startFromHead= false;
            count = lastItemIndex;
            current = this.tail
        }
        

        function resetCount(){
            if (startFromHead) {
                current = current?.getNext()
                count ++
            }else{
                current = current.getPrevious()
                count --
            }
        }

      
        while (current) {
            if (index === count) {
                return current
            }

            resetCount()
        }
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
    private previous?: LinkNode 
 
     constructor(private value: any){
     }
 
     setNext(nxt:LinkNode ){
         this.next = nxt;
         return this
     }
 
     getValue(){
         return this.value
     }
 
     getNext(){
         return this.next
     }

     getPrevious(){
        return this.previous
     }

     setPrevious(prev:LinkNode){
         this.previous =  prev

         return this
    }
 }

 let link =  new DoublyLinkedList("7748");
link.append("4")
.append(5)
.append(56)
.prepend("566")
.insert(3,"Miracle")

console.log( link.printList(),"\n\n",);

link.remove(2)

console.log( link.printList(),"\n\n",);

