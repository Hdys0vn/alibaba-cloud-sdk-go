package pairecservice

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

// ListLayers invokes the pairecservice.ListLayers API synchronously
func (client *Client) ListLayers(request *ListLayersRequest) (response *ListLayersResponse, err error) {
	response = CreateListLayersResponse()
	err = client.DoAction(request, response)
	return
}

// ListLayersWithChan invokes the pairecservice.ListLayers API asynchronously
func (client *Client) ListLayersWithChan(request *ListLayersRequest) (<-chan *ListLayersResponse, <-chan error) {
	responseChan := make(chan *ListLayersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLayers(request)
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

// ListLayersWithCallback invokes the pairecservice.ListLayers API asynchronously
func (client *Client) ListLayersWithCallback(request *ListLayersRequest, callback func(response *ListLayersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLayersResponse
		var err error
		defer close(result)
		response, err = client.ListLayers(request)
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

// ListLayersRequest is the request struct for api ListLayers
type ListLayersRequest struct {
	*requests.RoaRequest
	LaboratoryId string `position:"Query" name:"LaboratoryId"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

// ListLayersResponse is the response struct for api ListLayers
type ListLayersResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	TotalCount int64        `json:"TotalCount" xml:"TotalCount"`
	Layers     []LayersItem `json:"Layers" xml:"Layers"`
}

// CreateListLayersRequest creates a request to invoke ListLayers API
func CreateListLayersRequest() (request *ListLayersRequest) {
	request = &ListLayersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "ListLayers", "/api/v1/layers", "", "")
	request.Method = requests.GET
	return
}

// CreateListLayersResponse creates a response to parse from ListLayers response
func CreateListLayersResponse() (response *ListLayersResponse) {
	response = &ListLayersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
