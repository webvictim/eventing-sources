/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/webvictim/eventing-sources/pkg/kncloudevents"
)

// Example is a structure to hold CloudEvents
type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

/*
Example Output:

☁  cloudevents.Event:
Validation: valid
Context Attributes,
  SpecVersion: 0.2
  Type: dev.knative.eventing.samples.heartbeat
  Source: https://github.com/knative/eventing-sources/cmd/heartbeats/#local/demo
  ID: 3d2b5a1f-10ca-437b-a374-9c49e43c02fb
  Time: 2019-03-14T21:21:29.366002Z
  ContentType: application/json
  Extensions:
    the: 42
    beats: true
    heart: yes
Transport Context,
  URI: /
  Host: localhost:8080
  Method: POST
Data,
  {
    "id":162,
    "label":""
  }
*/

func display(event cloudevents.Event) {
	fmt.Printf("☁️  cloudevents.Event\n%s", event.String())
}

func gotEvent(event cloudevents.Event) {
	fmt.Printf("cloudevents.Event\n%s", event.String())

	fmt.Printf("Got Event Context: %+v\n", event.Context)

	var result map[string]interface{}
	dataBytes, err := event.DataBytes()
	if err != nil {
		json.Unmarshal(dataBytes, &result)

		fmt.Printf("Dumping result:")
		for key, value := range result {
			// Each value is an interface{} type, that is type asserted as a string
			fmt.Println(key, value.(string))
		}
	}

	fmt.Printf("----------------------------\n")
}

func main() {
	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), gotEvent))
}
