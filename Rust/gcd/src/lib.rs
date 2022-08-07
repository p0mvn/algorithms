
// gcd returns the greatest common divisor using the
// Euclidean algorithm.
// We would like to show that if a = bq + r -> gcd(a,b) = gcd(b,r)
// Proof:
// 
// a = bq + r
// gcd(a,b) = d iff gcd(b, r) = d
//
// 1. Prove gcd(a,b) = d -> gcd(b, r) = d
// Assume d | a and d | b
// Then, d| (a - qb) (can be proved separately but this is trivial)
// From a = bq + r, r = a - bq
// So, d | r
//
// 2. Prove gcd(b, r) = d ->  gcd(a,b) = d
// Asume d | b and d | r
// Then, d | bq + r
// That implies, d | a
// QED
fn gcd(a: u128, b: u128) -> u128 {
    if a > b {
        return gcd2(a, b);
    }
    return gcd2(b, a);
}

fn gcd2(a: u128, b: u128) -> u128 {
    let r = a % b;

    if r == 0 {
        return b;
    }

    return gcd(b, r);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn gcd_works() {
        let a = 2322;
        let b = 654;
        let expected_result = 6;

        let actual_result = gcd(a, b);

        assert_eq!(expected_result, actual_result);
    }

    #[test]
    fn gcd_reversed_works() {
        let a = 654;
        let b = 2322;
        let expected_result = 6;

        let actual_result = gcd(a, b);

        assert_eq!(expected_result, actual_result);
    }
}
