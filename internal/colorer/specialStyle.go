package colorer

func (colorer *Colorer) Reset() string {
	return "\033[0m"
}

func (colorer *Colorer) Bold() string {
	return "\033[1m"
}