class Solution {
public:
    int longestConsecutive(vector<int>& nums) {

        if (nums.empty()) {
            return 0;
        }
        
        unordered_set<int> set(nums.begin(), nums.end());
        int n = nums.size();
        int streak = 0;
        
        for (int num : set){
           if (set.find(num - 1) == set.end()) {
            int current_num = num;
            int current_streak = 1;

            while (set.find(current_num + 1) != set.end()) {
                current_num += 1;
                current_streak += 1;
            }

            streak = max(streak, current_streak);
           }
        }
        return streak;
    }
};
