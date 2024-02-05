package sprint

func Season(month string) string {

	switch month {
	case "dec", "jan", "feb":
		return "winter"
	case "mar", "apr", "may":
		return "spring"
	case "jun", "jul", "aug":
		return "summer"
	case "sep", "oct", "nov":
		return "autumn"
	default:
		return "invalid input: " + month
	}
}
