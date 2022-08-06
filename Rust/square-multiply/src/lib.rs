// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
// https://en.wikipedia.org/wiki/Modular_exponentiation

// exp_by_squaring_recusrsive_tail exponentiates base with exponent
// and returns the result using recursive tail method.
// Core idea:
//
// x^{n} = { (x^2)^{n / 2}              if even
//         { (x*x^2)^{(n - 1)/ 2}       if odd
// Note: the algorithm is non tail-recursive
fn exp_by_squaring_recursive_tail(base: u128, exponent: i128) -> u128 {
    // N.B.: to illustrate what could be done to suppport negative exponent.
    // Decimals or floats would have to be used for the result.
    // if exponent < 0 {
    //     return exp_by_squaring(1 / base, exponent * -1)
    // }
    if exponent == 0 {
        return 1
    }

    if exponent % 2 == 0 {
        return exp_by_squaring_recursive_tail(base * base, exponent / 2)
    }
    base * exp_by_squaring_recursive_tail(base * base, (exponent -1) / 2)
}


// exp_by_squaring_recusrsive_aux exponentiates base with exponent
// and returns the result using recursive accumulator method.
// Core idea:
//
// yx^{n} = { y(x^2)^{n / 2}              if even
//         { yx(x^2)^{(n - 1)/ 2}       if odd
// Note: the algorithm is non tail-recursive
fn exp_by_squaring_recursive_aux(base: u128, exponent: u128) -> u128 {
    exp_by_squaring_recursive_aux2(1, base, exponent)
}

fn exp_by_squaring_recursive_aux2(accum: u128, base: u128, exponent: u128) -> u128 {
    if exponent == 0 {
        return accum
    }

    if exponent % 2 == 0 {
        return exp_by_squaring_recursive_aux2(accum,  base * base, exponent / 2)
    }
    base * exp_by_squaring_recursive_aux2(accum * base, base * base, (exponent -1) / 2)
}

// exp_by_squaring_ierative exponentiates base with exponent
// and returns the result using iterative auxilary approach.
// It can be looked as transforming recusrsive aux to an iterative approach.
fn exp_by_squaring_iterative(mut base: u128, mut exponent: u128) -> u128 {
    let mut result = 1;
    while exponent > 0 {
        if exponent % 2 == 1 {
            result = result * base;
            exponent = exponent - 1;
        }
        base = base * base;
        exponent =  exponent  >> 1;
    }
    return result
}


// modular_exp computes modular exponentation by applying mod at each
// iteration step. The math gurantees that it is the same as if we were to
// apply mod at the end of the computation.
fn modular_exp(mut base: u128, mut exponent: u128, modulo: u128) -> u128 {
    let mut result = 1;
    while exponent > 0 {
        if exponent % 2 == 1 {
            result = result * base % modulo;
            exponent = exponent - 1;
        }
        base = base * base % modulo;
        exponent =  exponent  >> 1;
    }
    return result
} 

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn exp_by_squaring_recustive_tail_works() {
        let base = 3;
        let exponent = 45;

        let expected_result = 2954312706550833698643;

        let actual_result = exp_by_squaring_recursive_tail(base, exponent);
        assert_eq!(expected_result, actual_result);
    }

    #[test]
    fn exp_by_squaring_recustive_aux_works() {
        let base = 3;
        let exponent = 45;

        let expected_result = 2954312706550833698643;

        let actual_result = exp_by_squaring_recursive_tail(base, exponent);
        assert_eq!(expected_result, actual_result);
    }

    #[test]
    fn exp_by_squaring_iterative_works() {
        let base = 3;
        let exponent = 45;

        let expected_result = 2954312706550833698643;

        let actual_result = exp_by_squaring_iterative(base, exponent);
        assert_eq!(expected_result, actual_result);
    }

    #[test]
    fn modular_exp_works() {
        let base = 3;
        let exponent = 45;
        let modulo = 7;

        let expected_result =6;

        let actual_result = modular_exp(base, exponent, modulo);
        assert_eq!(expected_result, actual_result);
    }
}
