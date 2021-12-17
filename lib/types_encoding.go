// Code generated by fastssz. DO NOT EDIT.
// Hash: e2ec77411b97968084f5281f8c36a7216ec9a6e209a1e2ef84173fb5c244cd1f
package lib

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the TransactionsSSZ object
func (t *TransactionsSSZ) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(t)
}

// MarshalSSZTo ssz marshals the TransactionsSSZ object to a target array
func (t *TransactionsSSZ) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Transactions'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(t.Transactions); ii++ {
		offset += 4
		offset += len(t.Transactions[ii])
	}

	// Field (0) 'Transactions'
	if len(t.Transactions) > 1048576 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(t.Transactions)
		for ii := 0; ii < len(t.Transactions); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(t.Transactions[ii])
		}
	}
	for ii := 0; ii < len(t.Transactions); ii++ {
		if len(t.Transactions[ii]) > 1073741824 {
			err = ssz.ErrBytesLength
			return
		}
		dst = append(dst, t.Transactions[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the TransactionsSSZ object
func (t *TransactionsSSZ) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Transactions'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Transactions'
	{
		buf = tail[o0:]
		num, err := ssz.DecodeDynamicLength(buf, 1048576)
		if err != nil {
			return err
		}
		t.Transactions = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(t.Transactions[indx]) == 0 {
				t.Transactions[indx] = make([]byte, 0, len(buf))
			}
			t.Transactions[indx] = append(t.Transactions[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the TransactionsSSZ object
func (t *TransactionsSSZ) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Transactions'
	for ii := 0; ii < len(t.Transactions); ii++ {
		size += 4
		size += len(t.Transactions[ii])
	}

	return
}

// HashTreeRoot ssz hashes the TransactionsSSZ object
func (t *TransactionsSSZ) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(t)
}

// HashTreeRootWith ssz hashes the TransactionsSSZ object with a hasher
func (t *TransactionsSSZ) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Transactions'
	{
		subIndx := hh.Index()
		num := uint64(len(t.Transactions))
		if num > 1048576 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range t.Transactions {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1048576)
	}

	hh.Merkleize(indx)
	return
}
