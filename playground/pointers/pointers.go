package pointers

func GetAdultYears(age *int) {
	*age = *age - 18
}
