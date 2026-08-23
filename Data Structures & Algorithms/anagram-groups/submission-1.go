func groupAnagrams(strs []string) [][]string {
    mp := make(map[[26]byte][]string)

    for _, s := range strs {
        var count [26]byte
        for i := 0; i < len(s); i++ {
            count[s[i]-'a']++ 
        }
        mp[count] = append(mp[count], s)
    }

    result := make([][]string, 0, len(mp))
    for _, group := range mp {
        result = append(result, group)
    }
    return result
}
