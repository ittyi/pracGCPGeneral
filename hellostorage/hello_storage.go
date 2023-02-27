// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START functions_cloudevent_storage]

// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("HelloStorage", helloStorage)
}

// StorageObjectData contains metadata of the Cloud Storage object.
type StorageObjectData struct {
	Bucket         string    `json:"bucket,omitempty"`
	Name           string    `json:"name,omitempty"`
	Metageneration int64     `json:"metageneration,string,omitempty"`
	TimeCreated    time.Time `json:"timeCreated,omitempty"`
	Updated        time.Time `json:"updated,omitempty"`
	ContentType    string    `json:"contentType"`
}

// helloStorage consumes a CloudEvent message and logs details about the changed object.
func helloStorage(ctx context.Context, e event.Event) error {
	log.Printf("--helloStorage start--")
	helloStorageStart := time.Now()

	var data StorageObjectData
	if err := e.DataAs(&data); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}

	if strings.Contains(data.Name, ".json") == false {
		return nil
	}
	log.Printf("File: %s", data.Name)

	// // クライアントを作成する
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("NewClient errer")
		log.Fatal(err)
	}

	// Read the object1 from bucket.
	rc, err := client.Bucket(data.Bucket).Object(data.Name).NewReader(ctx)
	if err != nil {
		log.Printf("NewReader errer")
		log.Fatal(err)
	}

	// body, err := io.ReadAll(rc)
	// if err != nil {
	// 	log.Printf("io.ReadAll errer")
	// 	log.Fatal(err)
	// }

	// bodyStrList := strings.Split(string(body), "\n")
	// log.Printf("len: %d", len(bodyStrList))
	// for i := 0; i < len(bodyStrList); i++ {
	// 	log.Printf("body%d:%s", i, bodyStrList[i])
	// }

	scanner := bufio.NewScanner(rc)
	if err != nil {
		log.Printf("io.ReadAll errer")
		log.Fatal(err)
	}
	i := 1
	for scanner.Scan() {
		// ここで一行ずつ処理
		// fmt.Println(scanner.Text())
		log.Printf("body%d:%s", i, scanner.Text())
		i++
	}
	rc.Close()

	log.Printf("--helloStorage end--")
	helloStorageEnd := time.Now()
	log.Printf("difference:", helloStorageEnd.Sub(helloStorageStart))
	return nil
}

// [END functions_cloudevent_storage]
