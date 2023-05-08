package helper

func GreeterText(name string)string{
	if name==""{
		name ="World"
	}
	return "hello "+name
}