package colorer

func (colorer *Colorer) Green() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;0;255;0m"
	case "256":
		return "\033[38;5;46m"
	default:
		return "\033[32m"
	}
}

func (colorer *Colorer) Red() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;0;0m"
	case "256":
		return "\033[38;5;196m"
	default:
		return "\033[31m"
	}
}

func (colorer *Colorer) Cyan() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;0;255;255m"
	case "256":
		return "\033[38;5;51m"
	default:
		return "\033[36m"
	}
}
