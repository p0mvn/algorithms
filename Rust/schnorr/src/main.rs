use curve25519_dalek::EdwardsPoint;
use curve25519_dalek::Scalar;
use curve25519_dalek::constants::ED25519_BASEPOINT_POINT;

use rand::rngs::OsRng;

struct Prover {
    secret: Scalar,
}

impl Prover {
    fn commit(&self, generator: &EdwardsPoint) -> (Scalar, EdwardsPoint) {
        // Critical: k must be random. Otherwise, private key can be leaked if k is reused for two different
        // messages.
        let k = Scalar::random(&mut OsRng);
        let commitment = generator * k;
        (k, commitment)
    }

    fn respond(&self, k: Scalar, e: Scalar) -> Scalar {
        k + e * self.secret
    }
}

struct Verifier {
    public_key: EdwardsPoint,
}

impl Verifier {
    fn challenge() -> Scalar {
        Scalar::random(&mut OsRng)
    }
    fn verify(
        &self,
        generator: &EdwardsPoint,
        commitment: &EdwardsPoint,
        challenge: Scalar,
        response: Scalar,
    ) -> bool {
        // Schnorr signature verification works due to the algebraic construction of the response.
        // Here is why this check holds if the prover follows the protocol:
        //
        // Given:
        //   - Commitment R = [k]G          (G = generator, k = random nonce)
        //   - Public key    P = [x]G       (x = secret)
        //   - Challenge     e              (random scalar)
        //   - Response      z = k + e·x    (scalar addition, mod group order)
        //
        // Verifier checks:
        //   [z]G == R + [e]P
        //
        // Substitute values:
        //   Left:  [z]G         = [k + e·x]G   = [k]G + [e·x]G   = [k]G + [e]([x]G) = R + [e]P
        //   Right: R + [e]P
        //
        // So, if the signature (R, z) was honestly generated, both sides will be equal.
        //
        generator * response == commitment + challenge * self.public_key
    }
}

// 1. Pick random nonce k
// 2. Compute R = [k]G
// 3. Compute e = H(R || P || m)
// 4. Compute s = k + e·x  (mod n)
// 5. Return signature (R, s)

fn main() {
    // Create secret
    let secret = Scalar::from(1u64);

    // Define generator
    let generator: EdwardsPoint = ED25519_BASEPOINT_POINT;

    // Create public key from generator and secret
    let public_key: EdwardsPoint = generator * secret;

    // Initialize prover
    let prover = Prover { secret };

    // Initialize verifier
    let verifier: Verifier = Verifier { public_key };

    // Prover commits to a generator
    let (k, commitment) = prover.commit(&generator);

    // Verifier computes challenge
    let challenge1 = Verifier::challenge();

    // Prover responds to the challenge
    let response1: Scalar = prover.respond(k, challenge1);

    // Verifier verifies
    let is_verified = verifier.verify(&generator, &commitment, challenge1, response1);

    println!("{}", is_verified);

}
