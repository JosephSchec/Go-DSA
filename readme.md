Starting GoLang again while learning som DSA

### Things I learned about in Go

- When exporting in go functions and struct need to be capitalized
- How to index with array pointers
- Appending and popping from array pointers in functions
  - Appending is as simple as `arr = append(arr, value)`
  - Popping is done by pulling out the pointer referenc into a new variable, slicing it and reseteing the pointer

### Things I learned/relearned for DSA

- Big O explained (dropping constants)
- Linear/Binary Search
- Bubble Sort
- Deleting from linked list is constant time
- Similarities and Differences between Linked List and Arrays
  - Arrays are a set length regarless if you use the space or not while a linked list will only create new nodes when needed.
  - To access elements in a linked list you need to travers the whole thing vs array which you can index
- JS arrays are considered ArrayList with instant get/push/pop but have O(n) time for un/shift

### Things I learned/relearned about recursion

- Don't mix base case logic into the recusive function
  - I.e don't check if the curent value is out of bounds in the recursive function itself rather check that in the base case
- Establish the base cases first
  - A base case should be any time the code will fail or we don't want to continue (i.e in a maze if we go off the edge or we finish...)
- Recursion at least in the case of the maze is O(4n) 4 directions for each point and since we drop constants is O(n)
- When dealing with recursion you should add a `seen` property so in the case where you need to backtrack you don't get into an infinite loop
- 3 points in a recursion algorythm that you have access to the values after the base cases

  1. pre - here you can mark as seen and set add the current value to your stack
  2. recursion step - here you loop through checking base cases
  3. post - here you establish that if the last step in the recursive function didn't return true then we go to the previos recursive function and try a different route i.e in a maze that looks like this

  ```go
    maze := []string{
  	  "####S######",
  	  "####     ##",
  	  "####E######",
    }
  ```

  following the concept of going up->right->down->left means we will first go down once -> right->right->right (marking each point as seen as we go through it) then once we can't go further we will drop out of that last fn. Now we will try to go up but we can't(because there's a wall) then we will try right but we can't go (because that was marked as seen already) then down is a wall and we will be forced to go back again...recursivly.

- Recursion vs. For loop
  - For loop should be your first instinct
  - Anytime you deal with multiple branches or ways of traversing recursion should be on your mind

#### Resources

- Following along with [Frontend Masters class called "The Last Algorithms Course You'll need."](https://frontendmasters.com/courses/algorithms) and learning Golang along the way
