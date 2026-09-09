class Solution {
public:
    int maxArea(vector<int>& heights) {
        int left = 0;
        int right = heights.size() - 1;
        int max_area = 0;
        while (left < right) {
            int current_area = (right - left) * (min(heights[left],heights[right]));

            

            if (heights[left] < heights[right]){
                left++;
                max_area = max(max_area, current_area);
            }
            else if (heights[left] > heights[right]){
                right--;
                max_area = max(max_area, current_area);
            }
            else {
                max_area = max(max_area, current_area);
                left++;
                right--;
            }
        }

        return max_area;
    }
};
