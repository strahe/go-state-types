// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package eam

import (
	"fmt"
	"io"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

var lengthBufCreateParams = []byte{130}

func (t *CreateParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCreateParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Initcode ([]uint8) (slice)
	if len(t.Initcode) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Initcode was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Initcode))); err != nil {
		return err
	}

	if _, err := w.Write(t.Initcode[:]); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	return nil
}

func (t *CreateParams) UnmarshalCBOR(r io.Reader) error {
	*t = CreateParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Initcode ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Initcode: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Initcode = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Initcode[:]); err != nil {
		return err
	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	return nil
}

var lengthBufCreateReturn = []byte{131}

func (t *CreateReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCreateReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.ActorID (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.ActorID)); err != nil {
		return err
	}

	// t.RobustAddress (address.Address) (struct)
	if err := t.RobustAddress.MarshalCBOR(w); err != nil {
		return err
	}

	// t.EthAddress ([20]uint8) (array)
	if len(t.EthAddress) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.EthAddress was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.EthAddress))); err != nil {
		return err
	}

	if _, err := w.Write(t.EthAddress[:]); err != nil {
		return err
	}
	return nil
}

func (t *CreateReturn) UnmarshalCBOR(r io.Reader) error {
	*t = CreateReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ActorID (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ActorID = uint64(extra)

	}
	// t.RobustAddress (address.Address) (struct)

	{

		if err := t.RobustAddress.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.RobustAddress: %w", err)
		}

	}
	// t.EthAddress ([20]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.EthAddress: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 20 {
		return fmt.Errorf("expected array to have 20 elements")
	}

	t.EthAddress = [20]uint8{}

	if _, err := io.ReadFull(br, t.EthAddress[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufCreate2Params = []byte{130}

func (t *Create2Params) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCreate2Params); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Initcode ([]uint8) (slice)
	if len(t.Initcode) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Initcode was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Initcode))); err != nil {
		return err
	}

	if _, err := w.Write(t.Initcode[:]); err != nil {
		return err
	}

	// t.Salt ([32]uint8) (array)
	if len(t.Salt) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Salt was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Salt))); err != nil {
		return err
	}

	if _, err := w.Write(t.Salt[:]); err != nil {
		return err
	}
	return nil
}

func (t *Create2Params) UnmarshalCBOR(r io.Reader) error {
	*t = Create2Params{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Initcode ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Initcode: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Initcode = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Initcode[:]); err != nil {
		return err
	}
	// t.Salt ([32]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Salt: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 32 {
		return fmt.Errorf("expected array to have 32 elements")
	}

	t.Salt = [32]uint8{}

	if _, err := io.ReadFull(br, t.Salt[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufCreate2Return = []byte{131}

func (t *Create2Return) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCreate2Return); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.ActorID (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.ActorID)); err != nil {
		return err
	}

	// t.RobustAddress (address.Address) (struct)
	if err := t.RobustAddress.MarshalCBOR(w); err != nil {
		return err
	}

	// t.EthAddress ([20]uint8) (array)
	if len(t.EthAddress) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.EthAddress was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.EthAddress))); err != nil {
		return err
	}

	if _, err := w.Write(t.EthAddress[:]); err != nil {
		return err
	}
	return nil
}

func (t *Create2Return) UnmarshalCBOR(r io.Reader) error {
	*t = Create2Return{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ActorID (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ActorID = uint64(extra)

	}
	// t.RobustAddress (address.Address) (struct)

	{

		if err := t.RobustAddress.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.RobustAddress: %w", err)
		}

	}
	// t.EthAddress ([20]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.EthAddress: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 20 {
		return fmt.Errorf("expected array to have 20 elements")
	}

	t.EthAddress = [20]uint8{}

	if _, err := io.ReadFull(br, t.EthAddress[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufCreateAccount = []byte{129}

func (t *CreateAccount) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCreateAccount); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Pubkey ([65]uint8) (array)
	if len(t.Pubkey) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Pubkey was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Pubkey))); err != nil {
		return err
	}

	if _, err := w.Write(t.Pubkey[:]); err != nil {
		return err
	}
	return nil
}

func (t *CreateAccount) UnmarshalCBOR(r io.Reader) error {
	*t = CreateAccount{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Pubkey ([65]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Pubkey: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 65 {
		return fmt.Errorf("expected array to have 65 elements")
	}

	t.Pubkey = [65]uint8{}

	if _, err := io.ReadFull(br, t.Pubkey[:]); err != nil {
		return err
	}
	return nil
}