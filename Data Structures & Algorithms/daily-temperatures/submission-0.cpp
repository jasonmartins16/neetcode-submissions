class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        stack<int> temp;
        int n = temperatures.size();
        vector<int> result(n, 0);
        for (int i = 0; i < n; i++){
            while (!temp.empty() && temperatures[i] > temperatures[temp.top()]){
                int prev = temp.top();
                temp.pop();
                result[prev] = i - prev;
            }
            temp.push(i);
        }
        return result;
    }

};