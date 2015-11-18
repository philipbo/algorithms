# Sorting
---
## Bubble Sort
### Time Complexity:

- Best case: O(n)
- Average case: O(n^2)
- Worst case: O(n^2)

### Space Complexity:

- O(1)

## Insertion Sort

### Time Complexity:

- Best case: O(n) 
- Average case: O(n^2)
- Worst case: O(n^2)

### Space Complexity:

- O(n)


## Benchmark

    BenchmarkBubbleSort10-4    	20000000	        86.9 ns/op
    BenchmarkBubbleSort100-4   	  200000	      8225 ns/op
    BenchmarkBubbleSort1000-4  	    2000	    723684 ns/op
    BenchmarkBubbleSort10000-4 	      20	  76407577 ns/op
    BenchmarkBubbleSort100000-4	       1	22070055883 ns/op
    
    
    BenchmarkInsertionSort10-4    	100000000	        17.7 ns/op
    BenchmarkInsertionSort100-4   	10000000	       151 ns/op
    BenchmarkInsertionSort1000-4  	 1000000	      1379 ns/op
    BenchmarkInsertionSort10000-4 	  100000	     14143 ns/op
    BenchmarkInsertionSort100000-4	       1	3430866946 ns/op
