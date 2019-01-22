package strings


func Permute(s string) []string {

	if len(s) == 1 {
		a := make([]string,1)
		b := []rune(s)
		a[0] = string(b[0])
		return a
	}

	//fmt.Println(s)

	carr := []rune(s)
	letter := carr[0]
	l:= len(carr)

	str:= string(carr[1:l])

	res := Permute(str)

	var strarr []string
	

	strarr = make([]string,0)

	for i:=0;i<len(res);i++ {
		rarr := []rune(res[i])
		le := len(rarr)
		for j:=0 ;j<=len(rarr);j++{
			var sr string
			if j==0 {
				sr = string(letter) + string(rarr)
			}else{
				sr = string(rarr[0:j]) + string(letter) + string(rarr[j:le])
			}
			strarr = append(strarr,sr)
			//fmt.Println(sr)
	}
}



	return strarr


}