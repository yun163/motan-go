package serialize_test

import (
	"testing"
	"github.com/weibocom/motan-go/gengo/protodef"
	"github.com/weibocom/motan-go/serialize"
)

// serialize && deserialize string
func TestSerializeAddress(t *testing.T) {
	person := &tutorial.Person{}
	person.Id = 100
	person.Email = "wwwww"
	phoneNumber := new (tutorial.Person_PhoneNumber)
	person.Phones = append(person.Phones, phoneNumber)
	address := tutorial.AddressBook{}

	address.People = append(address.People, person)

	serialization := &serialize.ProtoBufSerialization{}

	bytes, err := serialization.Serialize(&address)

	if err != nil {
		t.Errorf("serialize address fail. err:%v\n", err)
	}
	address1 := new(tutorial.AddressBook)

	serialization.DeSerialize(bytes, address1)

	if address1 == nil {
		t.Errorf("Failed DeSerialize")
	} else {
		if len(address1.People) != 1 {
			t.Errorf("person count not equals 1")
		}
		person1 := address1.People[0]
		if person1.Email != "wwwww" {
			t.Errorf("Email invalid")
		}
	}
}
