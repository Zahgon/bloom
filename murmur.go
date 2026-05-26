/*
The bloom library relied on the excellent murmur library
by Sébastien Paolacci. Unfortunately, it involved some heap
allocation. We want to avoid any heap allocation whatsoever
in the hashing process. To preserve backward compatibility, we roll
our own hashing functions. They are designed to be strictly equivalent
to Paolacci's implementation.

License on original code:


Copyright 2013, Sébastien Paolacci.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the library nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

*/

package bloom

const (
	c1_128     = 0x87c37b91114253d5
	c2_128     = 0x4cf5ad432745937f
	block_size = 16
)

// digest128 represents a partial evaluation of a 128 bites hash.
type digest128 struct {
	h1 uint64 // Unfinalized running hash part 1.
	h2 uint64 // Unfinalized running hash part 2.
}

// bmix will hash blocks (16 bytes)
func (d *digest128) bmix(p []byte) { _ = "STUB: not implemented"; return }

// bmix_words will hash two 64-bit words (16 bytes)
func (d *digest128) bmix_words(k1, k2 uint64) { _ = "STUB: not implemented"; return }

// sum128 computers two 64-bit hash value. It is assumed that
// bmix was first called on the data to process complete blocks
// of 16 bytes. The 'tail' is a slice representing the 'tail' (leftover
// elements, fewer than 16). If pad_tail is true, we make it seem like
// there is an extra element with value 1 appended to the tail.
// The length parameter represents the full length of the data (including
// the blocks of 16 bytes, and, if pad_tail is true, an extra byte).
func (d *digest128) sum128(pad_tail bool, length uint, tail []byte) (h1, h2 uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func fmix64(k uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// sum256 will compute 4 64-bit hash values from the input.
// It is designed to never allocate memory on the heap. So it
// works without any byte buffer whatsoever.
// It is designed to be strictly equivalent to
//
//				a1 := []byte{1}
//	         hasher := murmur3.New128()
//	         hasher.Write(data) // #nosec
//	         v1, v2 := hasher.Sum128()
//	         hasher.Write(a1) // #nosec
//	         v3, v4 := hasher.Sum128()
//
// See TestHashRandom.
func (d *digest128) sum256(data []byte) (hash1, hash2, hash3, hash4 uint64) {
	_ = "STUB: not implemented"
	// We always start from zero.
	return 0, 0, 0, 0
}

// Process as many bytes as possible.

// We have enough to compute the first two 64-bit numbers

// Next we want to 'virtually' append 1 to the input, but,
// we do not want to append to an actual array!!!

// We are left with no tail!!!

// We append 1.

// We process the resulting 2 words.

// empty slice, deliberate.

// We still have a tail (fewer than 15 bytes) but we
// need to append '1' to it.
