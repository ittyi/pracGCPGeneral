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
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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
	log.Printf("Event ID: %s", e.ID())
	log.Printf("Event Type: %s", e.Type())
	log.Printf("DataMediaType: %s", e.DataMediaType())
	log.Printf("e.String(): %s", e.String())

	var data StorageObjectData
	if err := e.DataAs(&data); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}

	log.Printf("Bucket: %s", data.Bucket)
	log.Printf("File: %s", data.Name)
	log.Printf("Metageneration: %d", data.Metageneration)
	log.Printf("Created: %s", data.TimeCreated)
	log.Printf("Updated: %s", data.Updated)
	log.Printf("ContentType: %s", data.ContentType)

	s, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	log.Printf("string(s): %s", string(s))

	credentialFilePath := "../xxx.json"

	// クライアントを作成する
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	fmt.Println("test")
	fmt.Println(client)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	// GCSオブジェクトを書き込むファイルの作成
	f, err := os.Create("sample.txt")
	if err != nil {
		// log.Fatal(err)
	}

	fmt.Println(client)
	// オブジェクトのReaderを作成
	bucketName := "xxx"
	objectPath := "./xxx_2022_12_05_10_08_xxx.json"
	obj := client.Bucket(bucketName).Object(objectPath)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		fmt.Println("test")
		// log.Fatal(err)
	}
	defer reader.Close()

	// 書き込み
	tee := io.TeeReader(reader, f)
	s := bufio.NewScanner(tee)
	for s.Scan() {
	}
	if err := s.Err(); err != nil {
		// log.Fatal(err)
	}

	log.Println("done")

	return nil
}

// [END functions_cloudevent_storage]
