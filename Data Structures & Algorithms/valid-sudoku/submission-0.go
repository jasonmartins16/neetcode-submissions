func isValidSudoku(board [][]byte) bool {
    var rows [9][9]bool
    var cols [9][9]bool
    var boxes [9][9]bool

    for r := 0; r < 9; r++ {
        for c:= 0; c < 9; c++ {
            value := board[r][c]
            if value == '.' {
                continue
            }

            value -= '1'
            box_id := (r / 3) * 3 + (c / 3)

            if rows[r][value] || cols[c][value] || boxes[box_id][value] {
                return false
            }
            rows[r][value] = true
            cols[c][value] = true
            boxes[box_id][value] = true
        }
        
    }
    return true
}
