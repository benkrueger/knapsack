package common


func Max(x int,y int)(int){
	if(x > y){
		return x
	}else{
		return y
	}
}
func GCD(u int, v int)(int){
	for v != 0 {
		t := v
		v = u%v 
		u = t
	}
	return u
}
func VectorGCD(intlist []int)(int){
	resGCD := 0
	for _,i := range intlist {
		resGCD = GCD(resGCD,i)
	}
	return resGCD
}

/*
func kdp_solve(size int,l []*TorrentRecord)(int){
	num_entries := len(l)
	//var opt_list []TorrentEntry
	v := make([]int,num_entries)
	w := make([]int,num_entries)
	m := make([][]int,num_entries+1)
	
	for i,_ := range m {
		m[i]  = make([]int,size+1)
	}
	for i,t := range l {
		v[i] = t.Leeches
		w[i] = t.Length
	}
	//dynamic programming bit
	for i:= 1;i <= num_entries;i++ {
		for j:=0;j<=size;j++{
			if(i == 0 || j == 0){
				m[i][j] = 0
			}else if(w[i-1] <= j){
				m[i][j] = max(v[i-1]+m [i-1][j-w[i-1]],m[i-1][j])
			}else{
				m[i][j] = m[i-1][j]
			}
		}
	}

	return m[num_entries][size/v[0]]
}

*/
