class Solution {
public:

    string encode(vector<string>& strs) {
        string encode = "";
        for (const string& s : strs) {
            encode += to_string(s.size()) + "#" + s;
        }
        return encode;
    }

    vector<string> decode(string s) {
        vector<string> result;
        int i = 0;
        int n = s.size();

        while (i < n) {
            int j = i;
            while (s[j] != '#') {
                j++;
            }

            int len = stoi(s.substr(i, j - i));

            string word = s.substr(j + 1, len);
            result.push_back(word);

            i = j + 1 + len;
        }
        return result;
    }
};
