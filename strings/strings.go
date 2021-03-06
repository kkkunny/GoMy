package strings

// 去除两边空字符
func Strip(text string)string{
	for k, c := range text{
		if c != ' ' && c != '\n' && c != '\r' && c != '\t'{
			text = text[k:]
			break
		}
	}
	for i:=len(text)-1; i>=0; i--{
		c := text[i]
		if c != ' ' && c != '\n' && c != '\r' && c != '\t'{
			text = text[:i+1]
			break
		}
	}
	return text
}
// 获取字符串长度
func GetLength(data string)int{
	var length int
	for _, _ = range data{
		length++
	}
	return length
}