package merkletree

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

// SERIALIZE

func (b *Branch) Serialize(directory, serialiseTo string) {
	b.generateHash()

	switch serialiseTo {
	case "json":
		b.serializeJSON(directory)
	case "gob":
		b.serializeGob(directory)
	}

}

func (b *Branch) serializeJSON(directory string) {
	res, _ := json.Marshal(*b)

	f, err := os.Create(fmt.Sprintf("%s/%x.json", directory, b.Hash))
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(f)
	writer.Write(res)
	writer.Flush()
}

func (b *Branch) serializeGob(directory string) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)
	if err := enc.Encode(*b); err != nil {
		fmt.Println(err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%x", directory, b.Hash))
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(f)
	writer.Write(buf.Bytes())
	writer.Flush()
}

// DESERIALIZE
// ! TODO different types of serialized data

func Deserialize(filename string) (*Branch, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)
	b := make([]byte, reader.Size())

	_, err = reader.Read(b)
	if err != nil {
		return nil, err
	}

	// decode
	branch := Branch{}

	dec := gob.NewDecoder(bytes.NewBuffer(b))
	if err = dec.Decode(&branch); err != nil {
		return nil, err
	}

	fmt.Println(branch.LChild)

	return &branch, nil
}
