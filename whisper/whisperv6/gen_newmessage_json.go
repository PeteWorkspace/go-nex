// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package whisperv6

import (
	"encoding/json"

	"github.com/PeteWorkspace/go-nex/common/hexutil"
)

var _ = (*newMessageOverride)(nil)

// MarshalJSON marshals type NewMessage to a json string
func (n NewMessage) MarshalJSON() ([]byte, error) {
	type NewMessage struct {
		SymKeyID   string        `json:"symKeyID"`
		PublicKey  hexutil.Bytes `json:"pubKey"`
		Sig        string        `json:"sig"`
		TTL        uint32        `json:"ttl"`
		Topic      TopicType     `json:"topic"`
		Payload    hexutil.Bytes `json:"payload"`
		Padding    hexutil.Bytes `json:"padding"`
		PowTime    uint32        `json:"powTime"`
		PowTarget  float64       `json:"powTarget"`
		TargetPeer string        `json:"targetPeer"`
	}
	var enc NewMessage
	enc.SymKeyID = n.SymKeyID
	enc.PublicKey = n.PublicKey
	enc.Sig = n.Sig
	enc.TTL = n.TTL
	enc.Topic = n.Topic
	enc.Payload = n.Payload
	enc.Padding = n.Padding
	enc.PowTime = n.PowTime
	enc.PowTarget = n.PowTarget
	enc.TargetPeer = n.TargetPeer
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals type NewMessage to a json string
func (n *NewMessage) UnmarshalJSON(input []byte) error {
	type NewMessage struct {
		SymKeyID   *string        `json:"symKeyID"`
		PublicKey  *hexutil.Bytes `json:"pubKey"`
		Sig        *string        `json:"sig"`
		TTL        *uint32        `json:"ttl"`
		Topic      *TopicType     `json:"topic"`
		Payload    *hexutil.Bytes `json:"payload"`
		Padding    *hexutil.Bytes `json:"padding"`
		PowTime    *uint32        `json:"powTime"`
		PowTarget  *float64       `json:"powTarget"`
		TargetPeer *string        `json:"targetPeer"`
	}
	var dec NewMessage
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.SymKeyID != nil {
		n.SymKeyID = *dec.SymKeyID
	}
	if dec.PublicKey != nil {
		n.PublicKey = *dec.PublicKey
	}
	if dec.Sig != nil {
		n.Sig = *dec.Sig
	}
	if dec.TTL != nil {
		n.TTL = *dec.TTL
	}
	if dec.Topic != nil {
		n.Topic = *dec.Topic
	}
	if dec.Payload != nil {
		n.Payload = *dec.Payload
	}
	if dec.Padding != nil {
		n.Padding = *dec.Padding
	}
	if dec.PowTime != nil {
		n.PowTime = *dec.PowTime
	}
	if dec.PowTarget != nil {
		n.PowTarget = *dec.PowTarget
	}
	if dec.TargetPeer != nil {
		n.TargetPeer = *dec.TargetPeer
	}
	return nil
}
