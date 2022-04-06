class Graph {
    private vertices: number;
    private adjList: Record<number, number[]>;

    constructor(verticeNumber: number) {
        // V = v;
        this.adjList = {};
        this.vertices = verticeNumber;

        for (let i = 0; i < verticeNumber; ++i)
            this.adjList[i] = [];

    }

    public addEdge(v: number, w: number) {
        this.adjList[v].push(w);
    }

    // We are assuming that all nodes are connected for simplicity
    public BFS(vertice: number) {
        let queue = [vertice];
        let visited: Record<number, boolean> = {};
        visited[vertice] = true;

        while (queue.length) {
            let first = queue.shift();

            //Do the computation work if element is found
            console.log(first);


            for (let index = 0; index < this.adjList[first].length; index++) {
                const element = this.adjList[first][index];

                if (!visited[element]) {
                    queue.push(element);
                    visited[element] = true;
                }

            }

        }
    }

    // We are assuming that all nodes are connected for simplicity
    public DFS(vertice: number){
        let visited:Record<number, boolean> = {};
        // let 

        this.DFSutil(vertice, visited)
    }


    private DFSutil(vertice: number, visited:Record<number, boolean>){

        //Do the computation work if element is found
        console.log(vertice);
        visited[vertice] = true;

        for (let index = 0; index < this.adjList[vertice].length; index++) {
            const element = this.adjList[vertice][index];

            if (!visited[element]) {
                this.DFSutil(element, visited)
            }
        }


    }
}

let g = new Graph(4);

g.addEdge(0, 1);
g.addEdge(0, 2);
g.addEdge(1, 2);
g.addEdge(2, 0);
g.addEdge(2, 3);
g.addEdge(3, 3);

console.dir("Following is Breadth First Traversal "+
                           "(starting from vertex 2)");
 
        g.BFS(2);


console.dir("Following is Depth First Traversal "+
        "(starting from vertex 0)");

g.BFS(0);