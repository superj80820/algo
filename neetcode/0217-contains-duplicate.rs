use std::collections::HashSet;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut check = HashSet::new();
        for val in nums.iter() {
            if check.contains(val) {
                return true;
            } else {
                check.insert(val);
            }
        }
        return false;
    }
}
