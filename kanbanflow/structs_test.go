package kanbanflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true, "true should be true")
}

/*
 * Webhook event input data mapping tests
 */

func TestCreateTaskWebhook(t *testing.T) {
	data := []byte(`{
      "eventType": "taskCreated",
      "userId": "ec1b92fb1868c44aa9a041583c000e2a",
      "userFullName": "John Doe",
      "timestamp": "2015-10-20T14:45:06.331Z",
      "task": {
          "_id": "60e8b629fc8d6d28b513807d7d86b133",
          "name": "Write report",
          "description": "For school",
          "color": "green",
          "columnId": "ff31c6b2374911e49d115f7064763810",
          "totalSecondsSpent": 0,
          "totalSecondsEstimate": 0,
          "swimlaneId": "e037a6400e8911e5bdc9053860f3e5c0",
          "dates": [
              {
                  "targetColumnId": "ff31c6b4374911e49d115f7064763810",
                  "status": "active",
                  "dateType": "dueDate",
                  "dueTimestamp": "2015-10-20T15:00:00Z",
                  "dueTimestampLocal": "2015-10-20T17:00:00+02:00"
              }
          ],
          "subTasks": [
              {
                  "name": "Proofread",
                  "finished": false
              }
          ],
          "labels": [
              {
                  "name": "Writing",
                  "pinned": false
              }
          ]
      }
  }`)

	s := &CreateTaskWebhook{}
	encodeAndDecodeData(t, data, s)
}

func encodeAndDecodeData(t *testing.T, data []byte, s interface{}) {
	reader := bytes.NewReader(data)
	err := json.NewDecoder(reader).Decode(s)
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, err, "Should be able to decode data")
	buffer := bytes.NewBuffer([]byte{})
	err = json.NewEncoder(buffer).Encode(s)
	assert.NoError(t, err, "Should be able to encode")
	newData, err := ioutil.ReadAll(buffer)
	assert.NoError(t, err, "Shouldn't get an error from ioutil.ReadAll")

	assert.Equal(t,
		stripAllWhitespace(string(data)),
		stripAllWhitespace(string(newData)),
		"The data should be unchanged after decoding and encoding again",
	)

}

func stripAllWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
