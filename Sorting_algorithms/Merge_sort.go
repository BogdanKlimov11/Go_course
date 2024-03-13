package main

func merge(arr []int, left int, mid int, right int) {
  	n1 := mid - left + 1
  	n2 := right - mid
  
  	L := make([]int, n1)
  	R := make([]int, n2)
  
  	for i := 0; i < n1; i++ {
  		  L[i] = arr[left+i]
  	}
  	for j := 0; j < n2; j++ {
  		  R[j] = arr[mid+1+j]
  	}
  
  	i := 0
  	j := 0
  	k := left
  	for i < n1 && j < n2 {
    		if L[i] <= R[j] {
      			arr[k] = L[i]
      			i++
    		} else {
      			arr[k] = R[j]
      			j++
    		}
    		k++
  	}
  
  	for i < n1 {
    		arr[k] = L[i]
    		i++
    		k++
  	}
  
  	for j < n2 {
    		arr[k] = R[j]
    		j++
    		k++
  	}
}

func merge_sort(arr []int, left int, right int) {
    if left < right {
        mid := (left + right) / 2
        merge_sort(arr, left, mid)
        merge_sort(arr, mid+1, right)
        merge(arr, left, mid, right)
    }
}
