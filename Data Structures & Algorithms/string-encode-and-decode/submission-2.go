type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0
	n := len(encoded)

	for i < n {
		j := i
		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		end := start + length

		result = append(result, encoded[start:end])
		i = end
	}

	return result
}