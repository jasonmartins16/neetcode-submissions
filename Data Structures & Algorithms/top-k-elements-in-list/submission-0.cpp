class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        int n = nums.size();
        unordered_map<int, int> count;

        // count freq of each num
        for (int num : nums) {
            count[num]++;
        }

        // buckets arr where index are same as freq
        vector<vector<int>> buckets(n+1);
        for (const auto& entry : count) {
            int num = entry.first;
            int freq = entry.second;
            buckets[freq].push_back(num);
        }

        // collect elements from the highest freq to down to 1
        vector<int> result;
        result.reserve(k);
        for (int freq = n; freq >= 1; --freq) {
            for (int num : buckets[freq]) {
                result.push_back(num);
                if (result.size() == k) {
                    return result;
                }
            }
        }
        return result;
    }
    
};
