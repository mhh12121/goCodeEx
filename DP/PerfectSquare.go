package main

//bfs
//todo
func PerfectSquare(n int) int {
	perfectQ := make([]int, 0)
	cntPerfectQ := make([]int, n)
	for i := 1; i*i <= n; i++ {
		perfectQ = append(perfectQ, i*i)
		cntPerfectQ[i*i-1] = 1
	}
	if perfectQ[len(perfectQ)-1] == n {
		return 1
	}
	// Consider a graph which consists of number 0, 1,...,n as
	// its nodes. Node j is connected to node i via an edge if
	// and only if either j = i + (a perfect square number) or
	// i = j + (a perfect square number). Starting from node 0,
	// do the breadth-first search. If we reach node n at step
	// m, then the least number of perfect square numbers which
	// sum to n is m. Here since we have already obtained the
	// perfect square numbers, we have actually finished the
	// search at step 1.
	//
	searchQ := make([]int, 0)
	for i := 0; i < len(perfectQ); i++ {
		searchQ = append(searchQ, i)
	}

	currCntPerfectSquares := 1
	for len(searchQ) != 0 {
		currCntPerfectSquares++
		searchSize := len(searchQ)
		for i := 0; i < searchSize; i++ {
			tmp := searchQ[0]
			// Check the neighbors of node tmp which are the sum
			// of tmp and a perfect square number.
			for j := 0; j < len(perfectQ); j++ {
				if (tmp + j) == n {
					// We have reached node n
					return currCntPerfectSquares
				} else if ((tmp + j) < n) && cntPerfectQ[tmp+j-1] == 0 {
					// If cntPerfectSquares[tmp + j - 1] > 0, this is not
					// the first time that we visit this node and we should
					// skip the node (tmp + j).
					cntPerfectQ[tmp+j-1] = currCntPerfectSquares
					searchQ = append(searchQ, tmp+j)
				} else if (tmp + j) > n {
					// We don't need to consider the nodes which are greater than n.
					break
				}
			}
			searchQ = searchQ[:len(searchQ)-2]
		}

	}
	return 0
}
func PerfectSquare_dp(n int) int {
	if n <= 0 {
		return 0
	}
	cntPerfectQ := make([]int, n+1)
	for i := 0; i < len(cntPerfectQ); i++ {
		cntPerfectQ[i] = 99999999
	}
	cntPerfectQ[0] = 0
	for i := 1; i <= n; i++ {
		mint := 99999999
		j := 1
		for i-j*j >= 0 {
			mint = min(mint, cntPerfectQ[i-j*j]+1)
			j++
		}

		cntPerfectQ[i] = mint
	}
	return cntPerfectQ[n]
}

//lagrange's 4square
//todo
func PerfectSquare_lagra(n int) int {
	return 0
}
