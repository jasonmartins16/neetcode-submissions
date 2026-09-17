class Solution {
public:
    int evalRPN(vector<string>& tokens) {
       std::stack<int> st;
       int n = st.size();
       int result = 0;
       for (const std::string& token : tokens) {
            //if (st.empty()) {return result;}
            if (token != "+" && token != "-" && token != "*" && token != "/") {
                st.push(std::stoi(token));
            }
            else {
                int val2 = st.top(); st.pop();
                int val1 = st.top(); st.pop();
                if (token == "+") {
                    st.push(val1 + val2);
                    // Addition
                } else if (token == "-") {
                    st.push(val1 - val2);
                    // Subtraction
                } else if (token == "*") {
                    st.push(val1 * val2);
                    // Multiplication
                } else if (token == "/") {
                    st.push(val1 / val2);
                    // Division
                }
            }
       } 
       return st.top(); 
    }
};
