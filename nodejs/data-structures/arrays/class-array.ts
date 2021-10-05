class MyArray{
    length: number;
    data: Record<number, any>
    constructor(){
        this.length = 0;
        this.data = {}
    }

    get(index: number){
        return this.data[index]
    }

    push(item: any){
        this.data[this.length] = item;
        this.length ++

        return this.length;
    }

    pop(){
        if (this.length === 0) return;

       let lastItem =  this.data[this.length -1]
       delete this.data[this.length-1]

       return lastItem;
    }

    delete(index: number){
        if (this.length === 0) return;
        const item = this.data[index];
        this.shiftItems(index);
        this.length --
        return item
    }

    shiftItems(index: number){

        for (let ind = index; ind < this.length -1; ind++) {
            // const element = array[ind];
            this.data[ind] = this.data[ind+1]
        }
        
        delete this.data[this.length -1] 
    }
}

const theArray = new MyArray();

console.log(theArray.get(34), theArray);

theArray.push("hey")
theArray.push("nwanne")
theArray.push("telenovela")

console.log(theArray.delete(1));

console.log(theArray);


