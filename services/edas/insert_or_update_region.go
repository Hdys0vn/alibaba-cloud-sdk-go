package edas

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

// InsertOrUpdateRegion invokes the edas.InsertOrUpdateRegion API synchronously
func (client *Client) InsertOrUpdateRegion(request *InsertOrUpdateRegionRequest) (response *InsertOrUpdateRegionResponse, err error) {
	response = CreateInsertOrUpdateRegionResponse()
	err = client.DoAction(request, response)
	return
}

// InsertOrUpdateRegionWithChan invokes the edas.InsertOrUpdateRegion API asynchronously
func (client *Client) InsertOrUpdateRegionWithChan(request *InsertOrUpdateRegionRequest) (<-chan *InsertOrUpdateRegionResponse, <-chan error) {
	responseChan := make(chan *InsertOrUpdateRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertOrUpdateRegion(request)
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

// InsertOrUpdateRegionWithCallback invokes the edas.InsertOrUpdateRegion API asynchronously
func (client *Client) InsertOrUpdateRegionWithCallback(request *InsertOrUpdateRegionRequest, callback func(response *InsertOrUpdateRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertOrUpdateRegionResponse
		var err error
		defer close(result)
		response, err = client.InsertOrUpdateRegion(request)
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

// InsertOrUpdateRegionRequest is the request struct for api InsertOrUpdateRegion
type InsertOrUpdateRegionRequest struct {
	*requests.RoaRequest
	RegistryType  string           `position:"Query" name:"RegistryType"`
	DebugEnable   requests.Boolean `position:"Query" name:"DebugEnable"`
	RegionTag     string           `position:"Query" name:"RegionTag"`
	RegionName    string           `position:"Query" name:"RegionName"`
	Description   string           `position:"Query" name:"Description"`
	MseInstanceId string           `position:"Query" name:"MseInstanceId"`
	Id            requests.Integer `position:"Query" name:"Id"`
}

// InsertOrUpdateRegionResponse is the response struct for api InsertOrUpdateRegion
type InsertOrUpdateRegionResponse struct {
	*responses.BaseResponse
	Code                   int                    `json:"Code" xml:"Code"`
	Message                string                 `json:"Message" xml:"Message"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	UserDefineRegionEntity UserDefineRegionEntity `json:"UserDefineRegionEntity" xml:"UserDefineRegionEntity"`
}

// CreateInsertOrUpdateRegionRequest creates a request to invoke InsertOrUpdateRegion API
func CreateInsertOrUpdateRegionRequest() (request *InsertOrUpdateRegionRequest) {
	request = &InsertOrUpdateRegionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "InsertOrUpdateRegion", "/pop/v5/user_region_def", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInsertOrUpdateRegionResponse creates a response to parse from InsertOrUpdateRegion response
func CreateInsertOrUpdateRegionResponse() (response *InsertOrUpdateRegionResponse) {
	response = &InsertOrUpdateRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
