class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        hash1 = {}
        hash2 = {}
        for i in s:
            if i in hash1:
                hash1[i] += 1
            else:
                hash1[i] = 1
        
        for i in t:
            if i in hash2:
                hash2[i] += 1
            else:
                hash2[i] = 1

        if hash1 == hash2:
            return True
        else:
            return False       