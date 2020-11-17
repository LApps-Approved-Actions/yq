package yqlib

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/mikefarah/yq/v4/test"
)

var sampleYaml = `zabbix: winner
apple: great
banana:
- {cobra: kai, angus: bob}
`

var expectedJson = `{
  "zabbix": "winner",
  "apple": "great",
  "banana": [
    {
      "cobra": "kai",
      "angus": "bob"
    }
  ]
}
`

func TestJsonEncoderPreservesObjectOrder(t *testing.T) {
	var output bytes.Buffer
	writer := bufio.NewWriter(&output)

	var jsonEncoder = NewJsonEncoder(writer, 2)
	inputs, err := readDocuments(strings.NewReader(sampleYaml), "sample.yml")
	if err != nil {
		panic(err)
	}
	node := inputs.Front().Value.(*CandidateNode).Node
	jsonEncoder.Encode(node)
	writer.Flush()
	test.AssertResult(t, expectedJson, output.String())

}
