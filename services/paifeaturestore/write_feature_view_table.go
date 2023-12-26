package paifeaturestore

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// WriteFeatureViewTable invokes the paifeaturestore.WriteFeatureViewTable API synchronously
func (client *Client) WriteFeatureViewTable(request *WriteFeatureViewTableRequest) (response *WriteFeatureViewTableResponse, err error) {
	response = CreateWriteFeatureViewTableResponse()
	err = client.DoAction(request, response)
	return
}

// WriteFeatureViewTableWithChan invokes the paifeaturestore.WriteFeatureViewTable API asynchronously
func (client *Client) WriteFeatureViewTableWithChan(request *WriteFeatureViewTableRequest) (<-chan *WriteFeatureViewTableResponse, <-chan error) {
	responseChan := make(chan *WriteFeatureViewTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WriteFeatureViewTable(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// WriteFeatureViewTableWithCallback invokes the paifeaturestore.WriteFeatureViewTable API asynchronously
func (client *Client) WriteFeatureViewTableWithCallback(request *WriteFeatureViewTableRequest, callback func(response *WriteFeatureViewTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WriteFeatureViewTableResponse
		var err error
		defer close(result)
		response, err = client.WriteFeatureViewTable(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// WriteFeatureViewTableRequest is the request struct for api WriteFeatureViewTable
type WriteFeatureViewTableRequest struct {
	*requests.RoaRequest
	Body          string `position:"Body" name:"body"`
	InstanceId    string `position:"Path" name:"InstanceId"`
	FeatureViewId string `position:"Path" name:"FeatureViewId"`
}

// WriteFeatureViewTableResponse is the response struct for api WriteFeatureViewTable
type WriteFeatureViewTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateWriteFeatureViewTableRequest creates a request to invoke WriteFeatureViewTable API
func CreateWriteFeatureViewTableRequest() (request *WriteFeatureViewTableRequest) {
	request = &WriteFeatureViewTableRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "WriteFeatureViewTable", "/api/v1/instances/[InstanceId]/featureviews/[FeatureViewId]/action/writetable", "", "")
	request.Method = requests.POST
	return
}

// CreateWriteFeatureViewTableResponse creates a response to parse from WriteFeatureViewTable response
func CreateWriteFeatureViewTableResponse() (response *WriteFeatureViewTableResponse) {
	response = &WriteFeatureViewTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}