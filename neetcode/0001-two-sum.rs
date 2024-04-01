use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut check = HashMap::new();
        for (idx, val) in nums.iter().enumerate() {
            if let Some(&idx2) = check.get(&(target - val)) {
                return vec![idx as i32, idx2 as i32];
            } else {
                check.insert(val, idx as i32);
            }
        }
        unreachable!()
    }
}
