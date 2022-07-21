using System;
using System.Collections.Generic;
// you can also use other imports, for example:
// using System.Collections.Generic;

// you can write to stdout for debugging purposes, e.g.
// Console.WriteLine("this is a debug message");

class Solution {
    //create a list of arrays containg the crane ranges
    //sort the list
    //merge overlapping ranges
    //loop the merged ranges checking whether B and E belong to the same range
    //if in same range return true
    //else return false;
    public bool solution(int[] A, int[] P, int B, int E) {
        // write your code in C# 6.0 with .NET 4.5 (Mono)
        int len= A.Length;
        if(B==E){
            return true;
        }
        List<int[]> crane_ranges= new List<int[]>();

        for(int i=0; i<len; i++){
            crane_ranges.Add(new int[]{P[i]-A[i], P[i]+A[i]});            
        }
        crane_ranges.Sort((a,b)=>a[0]-b[0]);
      
        var mergedRanges= Merge(crane_ranges);       
 
        //check if B and E are in the same range
        int low=0;
        int high=mergedRanges.Count-1;
        while(low<=high){
            int mid= (low+high+1)/2;
            int [] arr= mergedRanges[mid];
            if(B>=arr[0]&& B<=arr[1]){
                if(E>=arr[0] && E<=arr[1]){
                    return true;
                }else{
                    return false;
                }
            }else if(B>=arr[0]){
                low=mid+1;
            }else{
                high=mid-1;
            }
        }
        return false;
    }
    public static List<int[]> Merge(List<int[]> intervals){
        var mergedRanges= new List<int[]>();
        int start=intervals[0][0];
        int end=intervals[0][1];

        foreach(var element in intervals){
            if(element[0]<=end){
                end = Math.Max(end, element[1]);

            }else{
                mergedRanges.Add(new int[] {(start), end});
                start=element[0];
                end=element[1];
            }
        }
        mergedRanges.Add(new int[]{start,end});
        return mergedRanges;
    }
   
}

// grinohdeveloper001