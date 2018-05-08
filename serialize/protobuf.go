package serialize

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"reflect"
)

type ProtoBufSerialization struct {
}

func (s *ProtoBufSerialization) GetSerialNum() int {
	return 7
}

func (s *ProtoBufSerialization) Serialize(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 100))
	if v == nil {
		buf.WriteByte(0)
		return buf.Bytes(), nil
	}
	var errRes error
	if nrv, ok := v.(proto.Message); ok {
	 	out, err := proto.Marshal(nrv)
		if err != nil {
			errRes = err
		} else {
			return out, nil
		}
	} else {
		fmt.Printf("Serialize Invalid type %s ", reflect.TypeOf(v).String())
		errRes = fmt.Errorf("can not Serialize. unknown type")
	}
	buf.WriteByte(0)
	return buf.Bytes(), errRes
}

func (s *ProtoBufSerialization) DeSerialize(b []byte, v interface{}) (interface{}, error) {
	if len(b) == 0 {
		return nil, nil
	}
	var errRes error
	if nrv, ok := v.(proto.Message); ok {
		if err := proto.Unmarshal(b, nrv); err != nil {
			errRes = fmt.Errorf("Failed to DeSerialize:", err)
		} else {
			return nrv, nil
		}
	} else {
		fmt.Printf("Invalid type %s ", reflect.TypeOf(v).String())
		errRes = fmt.Errorf("can not DeSerialize. unknown type")
	}
	return nil, errRes
}

func (s *ProtoBufSerialization) SerializeMulti(v []interface{}) ([]byte, error) {
	if len(v) != 1 {
		return nil, fmt.Errorf("Not support SerializeMulti")
	}
	return s.Serialize(v[0])
}

func (s *ProtoBufSerialization) DeSerializeMulti(b []byte, v []interface{}) (ret []interface{}, err error) {
	if len(v) != 1 {
		return nil, fmt.Errorf("Not support DeSerializeMulti")
	}
	if res, err := s.DeSerialize(b, v[0]); err == nil {
		return []interface{}{res}, nil
	} else {
		return nil, fmt.Errorf("Failed DeSerializeMulti for ", err)
	}
}