package chatbot

import "unicode"

// SplitCommandString - split command string
func SplitCommandString(str string) []string {
	var arr []string

	si := -1
	ysi := -1

	for i, v := range str {
		if ysi >= 0 {
			if v == '"' {
				if i > ysi+1 {
					arr = append(arr, str[(ysi+1):i])
				}

				ysi = -1
			}

			continue
		}

		if v == '"' {
			ysi = i
			si = -1
		} else if unicode.IsSpace(v) {
			if si < 0 {
				si = i + 1
			} else {
				if i == si {
					si = i + 1
				} else {
					arr = append(arr, str[si:i])

					si = -1
				}
			}
		} else {
			if si < 0 {
				si = i
			}
		}
	}

	if si >= 0 && si != len(str) {
		arr = append(arr, str[si:len(str)])
	}

	return arr
}

// SplitMultiCommandString - split multi command string
func SplitMultiCommandString(str string) []string {
	var arr []string

	si := -1
	ysi := -1

	for i, v := range str {
		if ysi >= 0 {
			if v == '"' {
				// if i > ysi+1 {
				// 	arr = append(arr, str[(ysi):i+1])
				// }

				ysi = -1
			}

			continue
		}

		if v == '"' {
			ysi = i
			// si = -1
		} else if v == '\n' || v == '\r' {
			if si < 0 {
				si = i + 1
			} else {
				if i == si {
					si = i + 1
				} else {
					arr = append(arr, str[si:i])

					si = -1
				}
			}
		} else {
			if si < 0 {
				si = i
			}
		}
	}

	if si >= 0 && si != len(str) {
		arr = append(arr, str[si:len(str)])
	}

	return arr
}
