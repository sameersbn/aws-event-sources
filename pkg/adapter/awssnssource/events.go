/*
Copyright (c) 2020 TriggerMesh Inc.

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

package awssnssource

import (
	"time"
)

// SNSEventRecord represents a SNS event.
type SNSEventRecord struct {
	EventVersion         string    `json:"eventVersion"`
	EventSubscriptionArn string    `json:"eventSubscriptionArn"`
	EventSource          string    `json:"eventSource"`
	SNS                  SNSEntity `json:"sns"`
}

// SNSEntity is the data from a SNS notification.
type SNSEntity struct {
	Signature         string                 `json:"signature"`
	MessageID         string                 `json:"messageId"`
	Type              string                 `json:"type"`
	TopicArn          string                 `json:"topicArn"`
	MessageAttributes map[string]interface{} `json:"messageAttributes"`
	SignatureVersion  string                 `json:"signatureVersion"`
	Timestamp         time.Time              `json:"timestamp"`
	SigningCertURL    string                 `json:"signingCertUrl"`
	Message           string                 `json:"message"`
	UnsubscribeURL    string                 `json:"unsubscribeUrl"`
	Subject           string                 `json:"subject"`
}
