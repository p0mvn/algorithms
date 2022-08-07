use std::cmp;

// can_measure_water retursn true if target_capacity can be reached by using the two jugs.
// ax + by = z, the greatest common divisor of a and b is c, z can also be divided by c.
// https://leetcode.com/problems/water-and-jug-problem
pub fn can_measure_water(jug1_capacity: i32, jug2_capacity: i32, target_capacity: i32) -> bool {
    let larger = cmp::max(jug1_capacity, jug2_capacity);
    let smaller = cmp::min(jug1_capacity, jug2_capacity);
    
    
    if target_capacity > larger + smaller {
        return false;
    }
    
    return target_capacity % gcd(larger, smaller) == 0;
}

pub fn gcd(a: i32, b: i32) -> i32 {
    if b == 0 {
        return a;
    }
    return gcd(b, a % b);
}
