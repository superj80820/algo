use std::collections::HashMap;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }
        let mut check = HashMap::new();
        for (s_char, t_char) in s.chars().zip(t.chars()) {
            *check.entry(s_char).or_insert(0) += 1;
            *check.entry(t_char).or_insert(0) -= 1;
        }
        check.into_values().all(|x| x == 0)
    }
}
