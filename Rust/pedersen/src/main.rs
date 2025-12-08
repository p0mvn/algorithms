use curve25519_dalek::constants::RISTRETTO_BASEPOINT_POINT;
use curve25519_dalek::ristretto::RistrettoPoint as Point;
use curve25519_dalek::scalar::Scalar;
use rand::rngs::OsRng;
use sha2::{Digest, Sha512};

// Generator points
fn generators() -> (Point, Point) {
    let G = RISTRETTO_BASEPOINT_POINT; // Standard base point

    // Derive H using "nothing up my sleeve" method
    let mut hasher = Sha512::new();
    hasher.update(b"Pedersen generator H");
    let hash = hasher.finalize();
    let hash_bytes: [u8; 64] = hash.into();
    let H = Point::from_uniform_bytes(&hash_bytes);

    (G, H)
}

// Commit: C = [v]G + [r]H
fn commit(v: Scalar, r: Scalar, G: Point, H: Point) -> Point {
    G * v + H * r
}

// Verify: check if C == [v]G + [r]H
fn verify(C: Point, v: Scalar, r: Scalar, G: Point, H: Point) -> bool {
    C == G * v + H * r
}

fn main() {
    let (G, H) = generators();

    let v1 = Scalar::from(1u64);
    let v2= Scalar::from(2u64);

    // Create random scalars for values and blinding factors
    let v1 = Scalar::from(100u64);  // Value 1
    let v2 = Scalar::from(250u64);  // Value 2
    let r1 = Scalar::random(&mut OsRng);  // Random blinding
    let r2 = Scalar::random(&mut OsRng);  // Random blinding
    
    // Commit to two values
    let C1 = commit(v1, r1, G, H);
    let C2 = commit(v2, r2, G, H);
    
    // Verify individual commitments
    assert!(verify(C1, v1, r1, G, H), "C1 should verify");
    assert!(verify(C2, v2, r2, G, H), "C2 should verify");
    
    // Test homomorphism: add commitments
    let C_sum = C1 + C2;

    // The sum opens to (v1 + v2) with blinding (r1 + r2)
    assert!(verify(C_sum, v1 + v2, r1 + r2, G, H), "Sum should verify!");
    
    // Test that wrong values fail
    let wrong_value = Scalar::from(999u64);
    assert!(!verify(C1, wrong_value, r1, G, H), "Wrong value should fail");
    
    println!("All tests passed!");
    println!("  - Individual commitments verify");
    println!("  - Homomorphic sum verifies: C1 + C2 opens to (v1+v2, r1+r2)");
    println!("  - Wrong values correctly rejected");
}
