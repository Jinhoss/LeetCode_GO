import re

class Solution:
    def myAtoi(self, s: str) -> int:
        INT_MIN, INT_MAX = -2**31, 2**31 - 1
        
        s = s.lstrip()
        match = re.match(r'^[+-]?\d+', s)
        if match:

            num_str = match.group(0)
            num = int(num_str)

            if num < INT_MIN:
                return INT_MIN
            if num > INT_MAX:
                return INT_MAX
            return num
        else:
            return 0
        