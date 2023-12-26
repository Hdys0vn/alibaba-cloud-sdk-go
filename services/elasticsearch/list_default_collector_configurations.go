package elasticsearch

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

// ListDefaultCollectorConfigurations invokes the elasticsearch.ListDefaultCollectorConfigurations API synchronously
func (client *Client) ListDefaultCollectorConfigurations(request *ListDefaultCollectorConfigurationsRequest) (response *ListDefaultCollectorConfigurationsResponse, err error) {
	response = CreateListDefaultCollectorConfigurationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDefaultCollectorConfigurationsWithChan invokes the elasticsearch.ListDefaultCollectorConfigurations API asynchronously
func (client *Client) ListDefaultCollectorConfigurationsWithChan(request *ListDefaultCollectorConfigurationsRequest) (<-chan *ListDefaultCollectorConfigurationsResponse, <-chan error) {
	responseChan := make(chan *ListDefaultCollectorConfigurationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDefaultCollectorConfigurations(request)
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

// ListDefaultCollectorConfigurationsWithCallback invokes the elasticsearch.ListDefaultCollectorConfigurations API asynchronously
func (client *Client) ListDefaultCollectorConfigurationsWithCallback(request *ListDefaultCollectorConfigurationsRequest, callback func(response *ListDefaultCollectorConfigurationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDefaultCollectorConfigurationsResponse
		var err error
		defer close(result)
		response, err = client.ListDefaultCollectorConfigurations(request)
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

// ListDefaultCollectorConfigurationsRequest is the request struct for api ListDefaultCollectorConfigurations
type ListDefaultCollectorConfigurationsRequest struct {
	*requests.RoaRequest
	ResType    string `position:"Query" name:"resType"`
	ResVersion string `position:"Query" name:"resVersion"`
	SourceType string `position:"Query" name:"sourceType"`
}

// ListDefaultCollectorConfigurationsResponse is the response struct for api ListDefaultCollectorConfigurations
type ListDefaultCollectorConfigurationsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListDefaultCollectorConfigurationsRequest creates a request to invoke ListDefaultCollectorConfigurations API
func CreateListDefaultCollectorConfigurationsRequest() (request *ListDefaultCollectorConfigurationsRequest) {
	request = &ListDefaultCollectorConfigurationsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListDefaultCollectorConfigurations", "/openapi/beats/default-configurations", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListDefaultCollectorConfigurationsResponse creates a response to parse from ListDefaultCollectorConfigurations response
func CreateListDefaultCollectorConfigurationsResponse() (response *ListDefaultCollectorConfigurationsResponse) {
	response = &ListDefaultCollectorConfigurationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}