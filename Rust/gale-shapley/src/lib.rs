use std::{error::Error};

// Algorithm definition:
// algorithm stable_matching is
//     Initialize m ∈ M and w ∈ W to free
//     while ∃ free man m who has a woman w to propose to do
//         w := first woman on m's list to whom m has not yet proposed
//         if ∃ some pair (m', w) then
//             if w prefers m to m' then
//                 m' becomes free
//                 (m, w) become engaged
//             end if
//         else
//             (m, w) become engaged
//         end if
//     repeat
//
// Runtime: O(n^2)

const UNMARRIED: i32 = -1;

// Performs gale shapley algorithm on the 2 lists of preferences to produce stable matching.
// Stability can be described as follows: for any matched pair, there is no case 
// where both a man and a woman prefer other partners.
//
// For every element i in men_pref, it is a preference of man i where preference is represented
// by ordering each woman j in women_pref
//
// For every element k in women_pref, it is a preference of woman k where preference is represented
// by ordering each man n in man_pref
//
// Assummes that all preferences are valid. That is, the length of men_pref is equal to the length of women_pref
// and point to existing persons in the opposite list.
//
// Returns matched pairs where the first element in the pair is the index of a man in men_pref
// and the second element is the index of a woman in women_pref.
pub fn gale_shapley(
    men_pref: & Vec<Vec<i32>>,
    women_pref: & Vec<Vec<i32>>,
) -> Result<Vec<(i32, i32)>, Box<dyn Error>> {
    if women_pref.len() != men_pref.len() {
        Err(format!(
            "men_pref length of {} does not equal women_pref length of {}",
            men_pref.len(),
            women_pref.len()
        ))?
    }

    // For woman i, represents the index of the man she is married to.
    // E.g. if married_pairs[1] == 3 means that woman 1 is married to man 3.
    //      if married_pairs[2] == -1 means that woman 2 is unmarried.
    let mut married_pairs: Vec<i32> = Vec::with_capacity(men_pref.len());

    // For man i, represents the index of the next woman in men_pref[i] preference list of man i.
    let mut men_next_choices: Vec<i32> = Vec::with_capacity(men_pref.len());

    let mut unmarried_men = Vec::with_capacity(men_pref.len());
    
    for i in 0..men_pref.len() {
        men_next_choices.push(0);
        married_pairs.push(UNMARRIED);
        unmarried_men.push(i as i32);
    }

    while unmarried_men.len() > 0 {
        let unmarried_man = unmarried_men[0];
        unmarried_men.drain(0..1);

        let cur_man_pref = &men_pref[unmarried_man as usize];

        let next_woman_idx = men_next_choices[unmarried_man as usize];
        
        // update the next choice in case gets unmarried by another man.
        men_next_choices[unmarried_man as usize] = next_woman_idx + 1;

        let next_woman_choice: i32 = cur_man_pref[next_woman_idx as usize];

        if married_pairs[next_woman_choice as usize] != UNMARRIED {
            // check if next_woman_choice prefers unmarried_man to her current fiance.
            let next_woman_pref = &women_pref[next_woman_choice as usize];
            for pref in next_woman_pref.iter() {
                if *pref == unmarried_man {
                    // prefers unmarried_man.

                    // the old fiance becomes unmarried.
                    unmarried_men.push(married_pairs[next_woman_choice as usize]);

                    // unmarried_man and next_woman_choice become engaged.
                    married_pairs[next_woman_choice as usize] = unmarried_man;
                    break
                } else if *pref == married_pairs[next_woman_choice as usize] {
                    // prefers her current choice.

                    // unmarried man keeps being unmarries, the old pair is preserved.
                    unmarried_men.push(unmarried_man);
                    break
                }
            }
        } else {
            // unmarried_man and next_woman_choice become engaged
            married_pairs[next_woman_choice as usize] = unmarried_man
        }
    }

    let mut result = Vec::with_capacity(married_pairs.len());

    for (woman, man) in married_pairs.iter().enumerate() {
        result.push((*man, woman as i32));
    }

    Ok(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn both_pref_empty() {
        let men_pref : Vec<Vec<i32>> = Vec::new();
        let women_pref: Vec<Vec<i32>> = Vec::new();

        let result = gale_shapley(&men_pref, &women_pref);

        let expected_result: Vec<(i32, i32)> = Vec::new();

        assert_eq!(expected_result, result.unwrap());
    }

    #[test]
    fn one_pair() {
        let men_pref : Vec<Vec<i32>> = vec![vec![0]];
        let women_pref : Vec<Vec<i32>> = vec![vec![0]];

        let result = gale_shapley(&men_pref, &women_pref);

        let expected_result: Vec<(i32, i32)> = vec![(0, 0)];

        assert_eq!(expected_result, result.unwrap());
    }

    #[test]
    fn two_pairs_pref_each_other() {
        let men_pref : Vec<Vec<i32>> = vec![vec![1, 0], vec![0, 1]];
        let women_pref : Vec<Vec<i32>> = vec![vec![1, 0], vec![0, 1]];

        let result = gale_shapley(&men_pref, &women_pref);

        let expected_result: Vec<(i32, i32)> = vec![(1, 0), (0, 1)];

        assert_eq!(expected_result, result.unwrap());
    }

    #[test]
    fn two_pairs_not_pref_each_other() {
        let men_pref : Vec<Vec<i32>> = vec![vec![1, 0], vec![0, 1]];
        let women_pref : Vec<Vec<i32>> = vec![vec![0, 1], vec![1, 0]];

        let result = gale_shapley(&men_pref, &women_pref);

        let expected_result: Vec<(i32, i32)> = vec![(1, 0), (0, 1)];

        assert_eq!(expected_result, result.unwrap());
    }

    #[test]
    fn four_pairs() {
        let men_pref : Vec<Vec<i32>> = vec![vec![0, 1, 3, 2], vec![0, 2, 3, 1], vec![1, 0, 2, 3], vec![0, 3, 1, 2]];
        let women_pref : Vec<Vec<i32>> = vec![vec![3, 1, 2, 0], vec![3, 1, 0, 2], vec![0, 3, 1, 2], vec![1, 0, 3, 2]];

        let result = gale_shapley(&men_pref, &women_pref);

        let expected_result: Vec<(i32, i32)> = vec![(3, 0), (0, 1), (1, 2), (2, 3)];

        assert_eq!(expected_result, result.unwrap());
    }

    #[test]
    #[should_panic(expected = "men_pref length of")]
    fn unequal_length_error() {
        let men_pref : Vec<Vec<i32>> = vec![vec![0, 1, 3, 2], vec![0, 2, 3, 1], vec![1, 0, 2, 3], ];
        let women_pref : Vec<Vec<i32>> = vec![vec![3, 1, 2, 0], vec![3, 1, 0, 2], vec![0, 3, 1, 2], vec![1, 0, 3, 2]];

        gale_shapley(&men_pref, &women_pref).unwrap();
    }
}
