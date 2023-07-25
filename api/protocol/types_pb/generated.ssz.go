// Code generated by fastssz. DO NOT EDIT.
// Hash: 4a603c93d7f3f3eaaa28c4e9f9aaf3faafda75d15184c2e608971ed88204de87
package types_pb

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the H128 object
func (h *H128) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H128 object to a target array
func (h *H128) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	dst = ssz.MarshalUint64(dst, h.Hi)

	// Field (1) 'Lo'
	dst = ssz.MarshalUint64(dst, h.Lo)

	return
}

// UnmarshalSSZ ssz unmarshals the H128 object
func (h *H128) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	h.Hi = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Lo'
	h.Lo = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H128 object
func (h *H128) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the H128 object
func (h *H128) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H128 object with a hasher
func (h *H128) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	hh.PutUint64(h.Hi)

	// Field (1) 'Lo'
	hh.PutUint64(h.Lo)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the H160 object
func (h *H160) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H160 object to a target array
func (h *H160) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H128)
	}
	if dst, err = h.Hi.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Lo'
	dst = ssz.MarshalUint32(dst, h.Lo)

	return
}

// UnmarshalSSZ ssz unmarshals the H160 object
func (h *H160) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 20 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H128)
	}
	if err = h.Hi.UnmarshalSSZ(buf[0:16]); err != nil {
		return err
	}

	// Field (1) 'Lo'
	h.Lo = ssz.UnmarshallUint32(buf[16:20])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H160 object
func (h *H160) SizeSSZ() (size int) {
	size = 20
	return
}

// HashTreeRoot ssz hashes the H160 object
func (h *H160) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H160 object with a hasher
func (h *H160) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	if err = h.Hi.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Lo'
	hh.PutUint32(h.Lo)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the H256 object
func (h *H256) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H256 object to a target array
func (h *H256) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H128)
	}
	if dst, err = h.Hi.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H128)
	}
	if dst, err = h.Lo.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the H256 object
func (h *H256) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 32 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H128)
	}
	if err = h.Hi.UnmarshalSSZ(buf[0:16]); err != nil {
		return err
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H128)
	}
	if err = h.Lo.UnmarshalSSZ(buf[16:32]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H256 object
func (h *H256) SizeSSZ() (size int) {
	size = 32
	return
}

// HashTreeRoot ssz hashes the H256 object
func (h *H256) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H256 object with a hasher
func (h *H256) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	if err = h.Hi.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Lo'
	if err = h.Lo.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the H384 object
func (h *H384) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H384 object to a target array
func (h *H384) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H256)
	}
	if dst, err = h.Hi.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H128)
	}
	if dst, err = h.Lo.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the H384 object
func (h *H384) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 48 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H256)
	}
	if err = h.Hi.UnmarshalSSZ(buf[0:32]); err != nil {
		return err
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H128)
	}
	if err = h.Lo.UnmarshalSSZ(buf[32:48]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H384 object
func (h *H384) SizeSSZ() (size int) {
	size = 48
	return
}

// HashTreeRoot ssz hashes the H384 object
func (h *H384) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H384 object with a hasher
func (h *H384) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	if err = h.Hi.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Lo'
	if err = h.Lo.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the H768 object
func (h *H768) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H768 object to a target array
func (h *H768) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H384)
	}
	if dst, err = h.Hi.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H384)
	}
	if dst, err = h.Lo.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the H768 object
func (h *H768) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 96 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H384)
	}
	if err = h.Hi.UnmarshalSSZ(buf[0:48]); err != nil {
		return err
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H384)
	}
	if err = h.Lo.UnmarshalSSZ(buf[48:96]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H768 object
func (h *H768) SizeSSZ() (size int) {
	size = 96
	return
}

// HashTreeRoot ssz hashes the H768 object
func (h *H768) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H768 object with a hasher
func (h *H768) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	if err = h.Hi.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Lo'
	if err = h.Lo.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the H512 object
func (h *H512) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H512 object to a target array
func (h *H512) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H256)
	}
	if dst, err = h.Hi.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H256)
	}
	if dst, err = h.Lo.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the H512 object
func (h *H512) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H256)
	}
	if err = h.Hi.UnmarshalSSZ(buf[0:32]); err != nil {
		return err
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H256)
	}
	if err = h.Lo.UnmarshalSSZ(buf[32:64]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H512 object
func (h *H512) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the H512 object
func (h *H512) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H512 object with a hasher
func (h *H512) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	if err = h.Hi.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Lo'
	if err = h.Lo.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the H1024 object
func (h *H1024) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H1024 object to a target array
func (h *H1024) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H512)
	}
	if dst, err = h.Hi.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H512)
	}
	if dst, err = h.Lo.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the H1024 object
func (h *H1024) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 128 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H512)
	}
	if err = h.Hi.UnmarshalSSZ(buf[0:64]); err != nil {
		return err
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H512)
	}
	if err = h.Lo.UnmarshalSSZ(buf[64:128]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H1024 object
func (h *H1024) SizeSSZ() (size int) {
	size = 128
	return
}

// HashTreeRoot ssz hashes the H1024 object
func (h *H1024) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H1024 object with a hasher
func (h *H1024) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	if err = h.Hi.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Lo'
	if err = h.Lo.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the H2048 object
func (h *H2048) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the H2048 object to a target array
func (h *H2048) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H1024)
	}
	if dst, err = h.Hi.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H1024)
	}
	if dst, err = h.Lo.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the H2048 object
func (h *H2048) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 256 {
		return ssz.ErrSize
	}

	// Field (0) 'Hi'
	if h.Hi == nil {
		h.Hi = new(H1024)
	}
	if err = h.Hi.UnmarshalSSZ(buf[0:128]); err != nil {
		return err
	}

	// Field (1) 'Lo'
	if h.Lo == nil {
		h.Lo = new(H1024)
	}
	if err = h.Lo.UnmarshalSSZ(buf[128:256]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the H2048 object
func (h *H2048) SizeSSZ() (size int) {
	size = 256
	return
}

// HashTreeRoot ssz hashes the H2048 object
func (h *H2048) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the H2048 object with a hasher
func (h *H2048) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hi'
	if err = h.Hi.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Lo'
	if err = h.Lo.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Block object
func (b *Block) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the Block object to a target array
func (b *Block) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'Header'
	dst = ssz.WriteOffset(dst, offset)
	if b.Header == nil {
		b.Header = new(Header)
	}
	offset += b.Header.SizeSSZ()

	// Offset (1) 'Body'
	dst = ssz.WriteOffset(dst, offset)
	if b.Body == nil {
		b.Body = new(Body)
	}
	offset += b.Body.SizeSSZ()

	// Field (0) 'Header'
	if dst, err = b.Header.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Block object
func (b *Block) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Header'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Body'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'Header'
	{
		buf = tail[o0:o1]
		if b.Header == nil {
			b.Header = new(Header)
		}
		if err = b.Header.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Body'
	{
		buf = tail[o1:]
		if b.Body == nil {
			b.Body = new(Body)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Block object
func (b *Block) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'Header'
	if b.Header == nil {
		b.Header = new(Header)
	}
	size += b.Header.SizeSSZ()

	// Field (1) 'Body'
	if b.Body == nil {
		b.Body = new(Body)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Block object
func (b *Block) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the Block object with a hasher
func (b *Block) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Header'
	if err = b.Header.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Header object
func (h *Header) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the Header object to a target array
func (h *Header) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(664)

	// Field (0) 'ParentHash'
	if h.ParentHash == nil {
		h.ParentHash = new(H256)
	}
	if dst, err = h.ParentHash.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Coinbase'
	if h.Coinbase == nil {
		h.Coinbase = new(H160)
	}
	if dst, err = h.Coinbase.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Root'
	if h.Root == nil {
		h.Root = new(H256)
	}
	if dst, err = h.Root.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'TxHash'
	if h.TxHash == nil {
		h.TxHash = new(H256)
	}
	if dst, err = h.TxHash.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'ReceiptHash'
	if h.ReceiptHash == nil {
		h.ReceiptHash = new(H256)
	}
	if dst, err = h.ReceiptHash.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'Difficulty'
	if h.Difficulty == nil {
		h.Difficulty = new(H256)
	}
	if dst, err = h.Difficulty.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (6) 'Number'
	if h.Number == nil {
		h.Number = new(H256)
	}
	if dst, err = h.Number.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, h.GasLimit)

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, h.GasUsed)

	// Field (9) 'Time'
	dst = ssz.MarshalUint64(dst, h.Time)

	// Field (10) 'Nonce'
	dst = ssz.MarshalUint64(dst, h.Nonce)

	// Field (11) 'BaseFee'
	if h.BaseFee == nil {
		h.BaseFee = new(H256)
	}
	if dst, err = h.BaseFee.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (12) 'Extra'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(h.Extra)

	// Field (13) 'Signature'
	if h.Signature == nil {
		h.Signature = new(H768)
	}
	if dst, err = h.Signature.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (14) 'Bloom'
	if h.Bloom == nil {
		h.Bloom = new(H2048)
	}
	if dst, err = h.Bloom.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (15) 'MixDigest'
	if h.MixDigest == nil {
		h.MixDigest = new(H256)
	}
	if dst, err = h.MixDigest.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (12) 'Extra'
	if size := len(h.Extra); size > 117 {
		err = ssz.ErrBytesLengthFn("--.Extra", size, 117)
		return
	}
	dst = append(dst, h.Extra...)

	return
}

// UnmarshalSSZ ssz unmarshals the Header object
func (h *Header) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 664 {
		return ssz.ErrSize
	}

	tail := buf
	var o12 uint64

	// Field (0) 'ParentHash'
	if h.ParentHash == nil {
		h.ParentHash = new(H256)
	}
	if err = h.ParentHash.UnmarshalSSZ(buf[0:32]); err != nil {
		return err
	}

	// Field (1) 'Coinbase'
	if h.Coinbase == nil {
		h.Coinbase = new(H160)
	}
	if err = h.Coinbase.UnmarshalSSZ(buf[32:52]); err != nil {
		return err
	}

	// Field (2) 'Root'
	if h.Root == nil {
		h.Root = new(H256)
	}
	if err = h.Root.UnmarshalSSZ(buf[52:84]); err != nil {
		return err
	}

	// Field (3) 'TxHash'
	if h.TxHash == nil {
		h.TxHash = new(H256)
	}
	if err = h.TxHash.UnmarshalSSZ(buf[84:116]); err != nil {
		return err
	}

	// Field (4) 'ReceiptHash'
	if h.ReceiptHash == nil {
		h.ReceiptHash = new(H256)
	}
	if err = h.ReceiptHash.UnmarshalSSZ(buf[116:148]); err != nil {
		return err
	}

	// Field (5) 'Difficulty'
	if h.Difficulty == nil {
		h.Difficulty = new(H256)
	}
	if err = h.Difficulty.UnmarshalSSZ(buf[148:180]); err != nil {
		return err
	}

	// Field (6) 'Number'
	if h.Number == nil {
		h.Number = new(H256)
	}
	if err = h.Number.UnmarshalSSZ(buf[180:212]); err != nil {
		return err
	}

	// Field (7) 'GasLimit'
	h.GasLimit = ssz.UnmarshallUint64(buf[212:220])

	// Field (8) 'GasUsed'
	h.GasUsed = ssz.UnmarshallUint64(buf[220:228])

	// Field (9) 'Time'
	h.Time = ssz.UnmarshallUint64(buf[228:236])

	// Field (10) 'Nonce'
	h.Nonce = ssz.UnmarshallUint64(buf[236:244])

	// Field (11) 'BaseFee'
	if h.BaseFee == nil {
		h.BaseFee = new(H256)
	}
	if err = h.BaseFee.UnmarshalSSZ(buf[244:276]); err != nil {
		return err
	}

	// Offset (12) 'Extra'
	if o12 = ssz.ReadOffset(buf[276:280]); o12 > size {
		return ssz.ErrOffset
	}

	if o12 < 664 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (13) 'Signature'
	if h.Signature == nil {
		h.Signature = new(H768)
	}
	if err = h.Signature.UnmarshalSSZ(buf[280:376]); err != nil {
		return err
	}

	// Field (14) 'Bloom'
	if h.Bloom == nil {
		h.Bloom = new(H2048)
	}
	if err = h.Bloom.UnmarshalSSZ(buf[376:632]); err != nil {
		return err
	}

	// Field (15) 'MixDigest'
	if h.MixDigest == nil {
		h.MixDigest = new(H256)
	}
	if err = h.MixDigest.UnmarshalSSZ(buf[632:664]); err != nil {
		return err
	}

	// Field (12) 'Extra'
	{
		buf = tail[o12:]
		if len(buf) > 117 {
			return ssz.ErrBytesLength
		}
		if cap(h.Extra) == 0 {
			h.Extra = make([]byte, 0, len(buf))
		}
		h.Extra = append(h.Extra, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Header object
func (h *Header) SizeSSZ() (size int) {
	size = 664

	// Field (12) 'Extra'
	size += len(h.Extra)

	return
}

// HashTreeRoot ssz hashes the Header object
func (h *Header) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the Header object with a hasher
func (h *Header) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	if err = h.ParentHash.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Coinbase'
	if err = h.Coinbase.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Root'
	if err = h.Root.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'TxHash'
	if err = h.TxHash.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'ReceiptHash'
	if err = h.ReceiptHash.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'Difficulty'
	if err = h.Difficulty.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (6) 'Number'
	if err = h.Number.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (7) 'GasLimit'
	hh.PutUint64(h.GasLimit)

	// Field (8) 'GasUsed'
	hh.PutUint64(h.GasUsed)

	// Field (9) 'Time'
	hh.PutUint64(h.Time)

	// Field (10) 'Nonce'
	hh.PutUint64(h.Nonce)

	// Field (11) 'BaseFee'
	if err = h.BaseFee.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (12) 'Extra'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(h.Extra))
		if byteLen > 117 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(h.Extra)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (117+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (117+31)/32)
		}
	}

	// Field (13) 'Signature'
	if err = h.Signature.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (14) 'Bloom'
	if err = h.Bloom.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (15) 'MixDigest'
	if err = h.MixDigest.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Verifier object
func (v *Verifier) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the Verifier object to a target array
func (v *Verifier) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'PublicKey'
	if v.PublicKey == nil {
		v.PublicKey = new(H384)
	}
	if dst, err = v.PublicKey.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Address'
	if v.Address == nil {
		v.Address = new(H160)
	}
	if dst, err = v.Address.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Verifier object
func (v *Verifier) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 68 {
		return ssz.ErrSize
	}

	// Field (0) 'PublicKey'
	if v.PublicKey == nil {
		v.PublicKey = new(H384)
	}
	if err = v.PublicKey.UnmarshalSSZ(buf[0:48]); err != nil {
		return err
	}

	// Field (1) 'Address'
	if v.Address == nil {
		v.Address = new(H160)
	}
	if err = v.Address.UnmarshalSSZ(buf[48:68]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Verifier object
func (v *Verifier) SizeSSZ() (size int) {
	size = 68
	return
}

// HashTreeRoot ssz hashes the Verifier object
func (v *Verifier) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the Verifier object with a hasher
func (v *Verifier) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'PublicKey'
	if err = v.PublicKey.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Address'
	if err = v.Address.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Reward object
func (r *Reward) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(r)
}

// MarshalSSZTo ssz marshals the Reward object to a target array
func (r *Reward) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Amount'
	if r.Amount == nil {
		r.Amount = new(H256)
	}
	if dst, err = r.Amount.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Address'
	if r.Address == nil {
		r.Address = new(H160)
	}
	if dst, err = r.Address.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Reward object
func (r *Reward) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 52 {
		return ssz.ErrSize
	}

	// Field (0) 'Amount'
	if r.Amount == nil {
		r.Amount = new(H256)
	}
	if err = r.Amount.UnmarshalSSZ(buf[0:32]); err != nil {
		return err
	}

	// Field (1) 'Address'
	if r.Address == nil {
		r.Address = new(H160)
	}
	if err = r.Address.UnmarshalSSZ(buf[32:52]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Reward object
func (r *Reward) SizeSSZ() (size int) {
	size = 52
	return
}

// HashTreeRoot ssz hashes the Reward object
func (r *Reward) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(r)
}

// HashTreeRootWith ssz hashes the Reward object with a hasher
func (r *Reward) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Amount'
	if err = r.Amount.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Address'
	if err = r.Address.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Body object
func (b *Body) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the Body object to a target array
func (b *Body) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'Txs'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Txs); ii++ {
		offset += 4
		offset += b.Txs[ii].SizeSSZ()
	}

	// Offset (1) 'Verifiers'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Verifiers) * 68

	// Offset (2) 'Rewards'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Rewards) * 52

	// Field (0) 'Txs'
	if size := len(b.Txs); size > 104857600 {
		err = ssz.ErrListTooBigFn("--.Txs", size, 104857600)
		return
	}
	{
		offset = 4 * len(b.Txs)
		for ii := 0; ii < len(b.Txs); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Txs[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Txs); ii++ {
		if dst, err = b.Txs[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (1) 'Verifiers'
	if size := len(b.Verifiers); size > 104857600 {
		err = ssz.ErrListTooBigFn("--.Verifiers", size, 104857600)
		return
	}
	for ii := 0; ii < len(b.Verifiers); ii++ {
		if dst, err = b.Verifiers[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (2) 'Rewards'
	if size := len(b.Rewards); size > 104857600 {
		err = ssz.ErrListTooBigFn("--.Rewards", size, 104857600)
		return
	}
	for ii := 0; ii < len(b.Rewards); ii++ {
		if dst, err = b.Rewards[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Body object
func (b *Body) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'Txs'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Verifiers'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'Rewards'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'Txs'
	{
		buf = tail[o0:o1]
		num, err := ssz.DecodeDynamicLength(buf, 104857600)
		if err != nil {
			return err
		}
		b.Txs = make([]*Transaction, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Txs[indx] == nil {
				b.Txs[indx] = new(Transaction)
			}
			if err = b.Txs[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (1) 'Verifiers'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 68, 104857600)
		if err != nil {
			return err
		}
		b.Verifiers = make([]*Verifier, num)
		for ii := 0; ii < num; ii++ {
			if b.Verifiers[ii] == nil {
				b.Verifiers[ii] = new(Verifier)
			}
			if err = b.Verifiers[ii].UnmarshalSSZ(buf[ii*68 : (ii+1)*68]); err != nil {
				return err
			}
		}
	}

	// Field (2) 'Rewards'
	{
		buf = tail[o2:]
		num, err := ssz.DivideInt2(len(buf), 52, 104857600)
		if err != nil {
			return err
		}
		b.Rewards = make([]*Reward, num)
		for ii := 0; ii < num; ii++ {
			if b.Rewards[ii] == nil {
				b.Rewards[ii] = new(Reward)
			}
			if err = b.Rewards[ii].UnmarshalSSZ(buf[ii*52 : (ii+1)*52]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Body object
func (b *Body) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'Txs'
	for ii := 0; ii < len(b.Txs); ii++ {
		size += 4
		size += b.Txs[ii].SizeSSZ()
	}

	// Field (1) 'Verifiers'
	size += len(b.Verifiers) * 68

	// Field (2) 'Rewards'
	size += len(b.Rewards) * 52

	return
}

// HashTreeRoot ssz hashes the Body object
func (b *Body) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the Body object with a hasher
func (b *Body) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Txs'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Txs))
		if num > 104857600 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Txs {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 104857600)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 104857600)
		}
	}

	// Field (1) 'Verifiers'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Verifiers))
		if num > 104857600 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Verifiers {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 104857600)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 104857600)
		}
	}

	// Field (2) 'Rewards'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Rewards))
		if num > 104857600 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Rewards {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 104857600)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 104857600)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Transaction object
func (t *Transaction) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(t)
}

// MarshalSSZTo ssz marshals the Transaction object to a target array
func (t *Transaction) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(336)

	// Field (0) 'Type'
	dst = ssz.MarshalUint64(dst, t.Type)

	// Field (1) 'Nonce'
	dst = ssz.MarshalUint64(dst, t.Nonce)

	// Field (2) 'GasPrice'
	if t.GasPrice == nil {
		t.GasPrice = new(H256)
	}
	if dst, err = t.GasPrice.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'Gas'
	dst = ssz.MarshalUint64(dst, t.Gas)

	// Field (4) 'FeePerGas'
	if t.FeePerGas == nil {
		t.FeePerGas = new(H256)
	}
	if dst, err = t.FeePerGas.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'PriorityFeePerGas'
	if t.PriorityFeePerGas == nil {
		t.PriorityFeePerGas = new(H256)
	}
	if dst, err = t.PriorityFeePerGas.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (6) 'Value'
	if t.Value == nil {
		t.Value = new(H256)
	}
	if dst, err = t.Value.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (7) 'Data'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(t.Data)

	// Offset (8) 'Sign'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(t.Sign)

	// Field (9) 'To'
	if t.To == nil {
		t.To = new(H160)
	}
	if dst, err = t.To.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (10) 'From'
	if t.From == nil {
		t.From = new(H160)
	}
	if dst, err = t.From.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (11) 'ChainID'
	dst = ssz.MarshalUint64(dst, t.ChainID)

	// Field (12) 'Hash'
	if t.Hash == nil {
		t.Hash = new(H256)
	}
	if dst, err = t.Hash.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (13) 'R'
	if t.R == nil {
		t.R = new(H256)
	}
	if dst, err = t.R.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (14) 'S'
	if t.S == nil {
		t.S = new(H256)
	}
	if dst, err = t.S.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (15) 'V'
	if t.V == nil {
		t.V = new(H256)
	}
	if dst, err = t.V.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (7) 'Data'
	if size := len(t.Data); size > 104857600 {
		err = ssz.ErrBytesLengthFn("--.Data", size, 104857600)
		return
	}
	dst = append(dst, t.Data...)

	// Field (8) 'Sign'
	if size := len(t.Sign); size > 104857600 {
		err = ssz.ErrBytesLengthFn("--.Sign", size, 104857600)
		return
	}
	dst = append(dst, t.Sign...)

	return
}

// UnmarshalSSZ ssz unmarshals the Transaction object
func (t *Transaction) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 336 {
		return ssz.ErrSize
	}

	tail := buf
	var o7, o8 uint64

	// Field (0) 'Type'
	t.Type = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Nonce'
	t.Nonce = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'GasPrice'
	if t.GasPrice == nil {
		t.GasPrice = new(H256)
	}
	if err = t.GasPrice.UnmarshalSSZ(buf[16:48]); err != nil {
		return err
	}

	// Field (3) 'Gas'
	t.Gas = ssz.UnmarshallUint64(buf[48:56])

	// Field (4) 'FeePerGas'
	if t.FeePerGas == nil {
		t.FeePerGas = new(H256)
	}
	if err = t.FeePerGas.UnmarshalSSZ(buf[56:88]); err != nil {
		return err
	}

	// Field (5) 'PriorityFeePerGas'
	if t.PriorityFeePerGas == nil {
		t.PriorityFeePerGas = new(H256)
	}
	if err = t.PriorityFeePerGas.UnmarshalSSZ(buf[88:120]); err != nil {
		return err
	}

	// Field (6) 'Value'
	if t.Value == nil {
		t.Value = new(H256)
	}
	if err = t.Value.UnmarshalSSZ(buf[120:152]); err != nil {
		return err
	}

	// Offset (7) 'Data'
	if o7 = ssz.ReadOffset(buf[152:156]); o7 > size {
		return ssz.ErrOffset
	}

	if o7 < 336 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (8) 'Sign'
	if o8 = ssz.ReadOffset(buf[156:160]); o8 > size || o7 > o8 {
		return ssz.ErrOffset
	}

	// Field (9) 'To'
	if t.To == nil {
		t.To = new(H160)
	}
	if err = t.To.UnmarshalSSZ(buf[160:180]); err != nil {
		return err
	}

	// Field (10) 'From'
	if t.From == nil {
		t.From = new(H160)
	}
	if err = t.From.UnmarshalSSZ(buf[180:200]); err != nil {
		return err
	}

	// Field (11) 'ChainID'
	t.ChainID = ssz.UnmarshallUint64(buf[200:208])

	// Field (12) 'Hash'
	if t.Hash == nil {
		t.Hash = new(H256)
	}
	if err = t.Hash.UnmarshalSSZ(buf[208:240]); err != nil {
		return err
	}

	// Field (13) 'R'
	if t.R == nil {
		t.R = new(H256)
	}
	if err = t.R.UnmarshalSSZ(buf[240:272]); err != nil {
		return err
	}

	// Field (14) 'S'
	if t.S == nil {
		t.S = new(H256)
	}
	if err = t.S.UnmarshalSSZ(buf[272:304]); err != nil {
		return err
	}

	// Field (15) 'V'
	if t.V == nil {
		t.V = new(H256)
	}
	if err = t.V.UnmarshalSSZ(buf[304:336]); err != nil {
		return err
	}

	// Field (7) 'Data'
	{
		buf = tail[o7:o8]
		if len(buf) > 104857600 {
			return ssz.ErrBytesLength
		}
		if cap(t.Data) == 0 {
			t.Data = make([]byte, 0, len(buf))
		}
		t.Data = append(t.Data, buf...)
	}

	// Field (8) 'Sign'
	{
		buf = tail[o8:]
		if len(buf) > 104857600 {
			return ssz.ErrBytesLength
		}
		if cap(t.Sign) == 0 {
			t.Sign = make([]byte, 0, len(buf))
		}
		t.Sign = append(t.Sign, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Transaction object
func (t *Transaction) SizeSSZ() (size int) {
	size = 336

	// Field (7) 'Data'
	size += len(t.Data)

	// Field (8) 'Sign'
	size += len(t.Sign)

	return
}

// HashTreeRoot ssz hashes the Transaction object
func (t *Transaction) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(t)
}

// HashTreeRootWith ssz hashes the Transaction object with a hasher
func (t *Transaction) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Type'
	hh.PutUint64(t.Type)

	// Field (1) 'Nonce'
	hh.PutUint64(t.Nonce)

	// Field (2) 'GasPrice'
	if err = t.GasPrice.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'Gas'
	hh.PutUint64(t.Gas)

	// Field (4) 'FeePerGas'
	if err = t.FeePerGas.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'PriorityFeePerGas'
	if err = t.PriorityFeePerGas.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (6) 'Value'
	if err = t.Value.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (7) 'Data'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(t.Data))
		if byteLen > 104857600 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(t.Data)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (104857600+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (104857600+31)/32)
		}
	}

	// Field (8) 'Sign'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(t.Sign))
		if byteLen > 104857600 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(t.Sign)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (104857600+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (104857600+31)/32)
		}
	}

	// Field (9) 'To'
	if err = t.To.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (10) 'From'
	if err = t.From.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (11) 'ChainID'
	hh.PutUint64(t.ChainID)

	// Field (12) 'Hash'
	if err = t.Hash.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (13) 'R'
	if err = t.R.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (14) 'S'
	if err = t.S.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (15) 'V'
	if err = t.V.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
