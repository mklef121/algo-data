class HashTable {
    data: any[]
    constructor(size: number){
        this.data = new Array(size)
    }

   private _hash(key: string){
        let hash = 0
        for (let index = 0; index < key.length; index++) {
            // const element = array[index];
            //return a random number that is withing the range of the size of the hash array
            hash = (hash + key.charCodeAt(index)*index)%this.data.length
        }

        return hash
    }

    set(key:string, value){
        let address = this._hash(key);

        if (!this.data[address])  this.data[address] = []

        this.data[address].push([key, value])
    }

    get(key: string){
        let address = this._hash(key);
        let data = this.data[address]
        if (!data) return

        for (let index = 0; index < data.length; index++) {
            if (data[index][0] == key) {
                return data[index][1]
            }
        }

        return
    }

    keys() {
        if (!this.data.length) {
          return undefined
        }
        let result = []
        // loop through all the elements
        for (let i = 0; i < this.data.length; i++) {
            // if it's not an empty memory cell
            if (this.data[i] && this.data[i].length) {
              // but also loop through all the potential collisions
              if (this.data.length > 1) {
                for (let j = 0; j < this.data[i].length; j++) {
                  result.push(this.data[i][j][0])
                }
              } else {
                result.push(this.data[i][0])
              } 
            }
        }
        return result; 
      }
}

const hash = new HashTable(2);

hash.set("grapes", 1000)
hash.set("apples", 1000)

console.log(hash.get('grapes'));

console.log(hash.keys())
