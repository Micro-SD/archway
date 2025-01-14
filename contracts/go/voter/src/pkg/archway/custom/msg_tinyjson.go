// Code generated by tinyjson for marshaling/unmarshaling. DO NOT EDIT.

package custom

import (
	types "github.com/CosmWasm/cosmwasm-go/std/types"
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

func tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom(in *jlexer.Lexer, out *WithdrawRewardsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "records_num":
			out.RecordsNum = uint64(in.Uint64())
		case "total_rewards":
			if in.IsNull() {
				in.Skip()
				out.TotalRewards = nil
			} else {
				in.Delim('[')
				if out.TotalRewards == nil {
					if !in.IsDelim(']') {
						out.TotalRewards = make([]types.Coin, 0, 2)
					} else {
						out.TotalRewards = []types.Coin{}
					}
				} else {
					out.TotalRewards = (out.TotalRewards)[:0]
				}
				for !in.IsDelim(']') {
					var v1 types.Coin
					(v1).UnmarshalTinyJSON(in)
					out.TotalRewards = append(out.TotalRewards, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom(out *jwriter.Writer, in WithdrawRewardsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"records_num\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.RecordsNum))
	}
	{
		const prefix string = ",\"total_rewards\":"
		out.RawString(prefix)
		if in.TotalRewards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.TotalRewards {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalTinyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WithdrawRewardsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v WithdrawRewardsResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WithdrawRewardsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *WithdrawRewardsResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom(l, v)
}
func tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom1(in *jlexer.Lexer, out *WithdrawRewardsRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "records_limit":
			if in.IsNull() {
				in.Skip()
				out.RecordsLimit = nil
			} else {
				if out.RecordsLimit == nil {
					out.RecordsLimit = new(uint64)
				}
				*out.RecordsLimit = uint64(in.Uint64())
			}
		case "record_ids":
			if in.IsNull() {
				in.Skip()
				out.RecordIds = nil
			} else {
				in.Delim('[')
				if out.RecordIds == nil {
					if !in.IsDelim(']') {
						out.RecordIds = make([]uint64, 0, 8)
					} else {
						out.RecordIds = []uint64{}
					}
				} else {
					out.RecordIds = (out.RecordIds)[:0]
				}
				for !in.IsDelim(']') {
					var v4 uint64
					v4 = uint64(in.Uint64())
					out.RecordIds = append(out.RecordIds, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom1(out *jwriter.Writer, in WithdrawRewardsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"records_limit\":"
		out.RawString(prefix[1:])
		if in.RecordsLimit == nil {
			out.RawString("null")
		} else {
			out.Uint64(uint64(*in.RecordsLimit))
		}
	}
	{
		const prefix string = ",\"record_ids\":"
		out.RawString(prefix)
		if in.RecordIds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.RecordIds {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WithdrawRewardsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v WithdrawRewardsRequest) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WithdrawRewardsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom1(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *WithdrawRewardsRequest) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom1(l, v)
}
func tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom2(in *jlexer.Lexer, out *UpdateMetadataRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "owner_address":
			out.OwnerAddress = string(in.String())
		case "rewards_address":
			out.RewardsAddress = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom2(out *jwriter.Writer, in UpdateMetadataRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"owner_address\":"
		out.RawString(prefix[1:])
		out.String(string(in.OwnerAddress))
	}
	{
		const prefix string = ",\"rewards_address\":"
		out.RawString(prefix)
		out.String(string(in.RewardsAddress))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateMetadataRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v UpdateMetadataRequest) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateMetadataRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom2(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *UpdateMetadataRequest) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom2(l, v)
}
func tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom3(in *jlexer.Lexer, out *RewardsMsg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "update_metadata":
			if in.IsNull() {
				in.Skip()
				out.UpdateMetadata = nil
			} else {
				if out.UpdateMetadata == nil {
					out.UpdateMetadata = new(UpdateMetadataRequest)
				}
				(*out.UpdateMetadata).UnmarshalTinyJSON(in)
			}
		case "withdraw_rewards":
			if in.IsNull() {
				in.Skip()
				out.WithdrawRewards = nil
			} else {
				if out.WithdrawRewards == nil {
					out.WithdrawRewards = new(WithdrawRewardsRequest)
				}
				(*out.WithdrawRewards).UnmarshalTinyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom3(out *jwriter.Writer, in RewardsMsg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"update_metadata\":"
		out.RawString(prefix[1:])
		if in.UpdateMetadata == nil {
			out.RawString("null")
		} else {
			(*in.UpdateMetadata).MarshalTinyJSON(out)
		}
	}
	{
		const prefix string = ",\"withdraw_rewards\":"
		out.RawString(prefix)
		if in.WithdrawRewards == nil {
			out.RawString("null")
		} else {
			(*in.WithdrawRewards).MarshalTinyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RewardsMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v RewardsMsg) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RewardsMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom3(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *RewardsMsg) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom3(l, v)
}
func tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom4(in *jlexer.Lexer, out *CustomMsg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "rewards":
			if in.IsNull() {
				in.Skip()
				out.Rewards = nil
			} else {
				if out.Rewards == nil {
					out.Rewards = new(RewardsMsg)
				}
				(*out.Rewards).UnmarshalTinyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom4(out *jwriter.Writer, in CustomMsg) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Rewards != nil {
		const prefix string = ",\"rewards\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Rewards).MarshalTinyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v CustomMsg) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjsonF5cd6cf9EncodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom4(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *CustomMsg) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjsonF5cd6cf9DecodeGithubComArchwayNetworkVoterSrcPkgArchwayCustom4(l, v)
}
