package shape
type Rectangel struct{
Width float32
Length float32
}
func (rectangel*Rectangel) Area() float32{
	return rectangel.Width * rectangel.Length
}
func (rectangel *Rectangel) Perimeter() float32{
	return 2* (rectangel.Width+rectangel.Length)
}



