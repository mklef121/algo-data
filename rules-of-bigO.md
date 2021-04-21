
#### Rules to define the Complexity of BigO

There are several rules to find how the BigO of a function or a code since several scenarios might be playing out
in the said function.


- Rule 1: **The Worst Case**

Assume that the function operates at it's worst scenario. E.G

```go
package main;

hold:= []string{
    "hi","come","Nemo","Dull","Rahman","Django","Tall","Build"
}

func main(){
    for  i:=0; i < len(hold); i++ {
        if hold[i] == "Nemo"{
            break;
        }
    }
}

//if this main function runs, we can say it's Big O(3), but worst case scenario will assume the array is larger
// Than what it is currently, and that `Nemo` is the last element of the array
```
Thus we will assume that this is Big O(n)