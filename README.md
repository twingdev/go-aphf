# Homomorphic Digital Signatures Package based on Asynchroneous Programmable Hash Functions (APHF)

This package provides a robust implementation of digital signatures using elliptic curve cryptography (ECC) with the BN256 curve and the BLAKE3 hashing algorithm. It includes functionality for generating keys, signing data, and verifying signatures, all while utilizing a programmable hash function influenced by a trapdoor mechanism.

## Features

- **Elliptic Curve Cryptography**: Uses the BN256 curve for secure digital signatures.
- **BLAKE3 Hashing**: Integrates the BLAKE3 cryptographic hash function for fast and secure data hashing.
- **Trapdoor Influence**: Demonstrates the use of a trapdoor to alter the hash function behavior in signatures.
- **Hex Encoding**: Utilizes hexadecimal encoding for efficient string representation of signatures.

## Inspiration

This implementation was inspired by the concepts presented in the paper ["Programmable Hash Functions Go Private: Constructions and Applications to (Homomorphic) Signatures with Shorter Public Keys"](https://link.springer.com/chapter/10.1007/978-3-662-48000-7_13) by Dario Catalano, Dario Fiore, and Luca Nizzardo. The paper discusses programmable hash functions and their applications in cryptography, particularly in creating more efficient and secure digital signatures. The ideas from this research have guided the architectural choices and features of this digital signatures package.

## Concept

The concept of an Asynchronous Programmable Hash Function (APHF) builds on the idea of programmable hash functions, extending it into a setting where hash function behavior can be asymmetrically influenced by different parties or under different circumstances. Although "asynchronous" might not be the standard term, the interpretation might focus on how different aspects of the system (such as different computational or operational phases) can independently contribute to the overall behavior of the hash function.

**Background: Programmable Hash Functions (PHFs)**

To better understand APHFs, it’s essential to first look at the concept of Programmable Hash Functions (PHFs), which were introduced in the cryptographic literature to emulate random oracles in a programmable way. A PHF allows a designer to influence the hash function's output in a way that is indistinguishable from random by any observer not privy to the secret (programmable) settings.

**Asymmetric Programmable Hash Functions (APHFs)**

In an APHF scenario, you extend the idea of PHFs into a domain where the programmability or the output can be controlled or influenced asymmetrically. This can mean:

Asymmetric Information: Different parties might have different levels of influence or visibility into how the hash function is programmed. For instance, one party might be able to set conditions under which certain outputs are more likely, while another might only see the resulting hash values without understanding the underlying conditions.


Temporal Asymmetry: Different parameters might be used to program the hash function at different times, affecting its behavior in stages. This could be useful in protocols where phases (like setup, operation, and teardown) need different security properties from the hash function.
Potential Applications

Homomorphic Signatures: APHFs can be designed to support homomorphic properties, allowing computations on hashed data that still retain the hash's validity under certain algebraic operations. This is particularly useful in scenarios like cloud computing, where data integrity and computation verification need to be maintained without revealing the actual data.

Conditional Privacy: APHFs can be employed to provide privacy-preserving functionalities where hash outputs reveal information conditionally based on additional inputs (like keys or other cryptographic elements) that might be revealed later. This could be used in voting systems, where vote integrity is maintained without revealing voter choices until a specific condition (like the voting period ending) is met.

Cryptographic Commitments: Leveraging APHFs in commitment schemes where the commitment's reveal phase is programmable by an initial setup that remains hidden until the reveal phase. This could be used in scenarios involving timed releases or staged disclosures.


## Installation

To install and use this digital signatures package, follow these steps:

### Prerequisites

Ensure you have Go installed on your system. This package requires Go 1.15 or later.

### Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/genovatix/go-aphf.git
