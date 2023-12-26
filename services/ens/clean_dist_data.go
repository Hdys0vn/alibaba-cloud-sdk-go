package ens

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

// CleanDistData invokes the ens.CleanDistData API synchronously
func (client *Client) CleanDistData(request *CleanDistDataRequest) (response *CleanDistDataResponse, err error) {
	response = CreateCleanDistDataResponse()
	err = client.DoAction(request, response)
	return
}

// CleanDistDataWithChan invokes the ens.CleanDistData API asynchronously
func (client *Client) CleanDistDataWithChan(request *CleanDistDataRequest) (<-chan *CleanDistDataResponse, <-chan error) {
	responseChan := make(chan *CleanDistDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CleanDistData(request)
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

// CleanDistDataWithCallback invokes the ens.CleanDistData API asynchronously
func (client *Client) CleanDistDataWithCallback(request *CleanDistDataRequest, callback func(response *CleanDistDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CleanDistDataResponse
		var err error
		defer close(result)
		response, err = client.CleanDistData(request)
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

// CleanDistDataRequest is the request struct for api CleanDistData
type CleanDistDataRequest struct {
	*requests.RpcRequest
	EnsRegionId string `position:"Query" name:"EnsRegionId"`
	DataName    string `position:"Query" name:"DataName"`
	DataVersion string `position:"Query" name:"DataVersion"`
	AppId       string `position:"Query" name:"AppId"`
}

// CleanDistDataResponse is the response struct for api CleanDistData
type CleanDistDataResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCleanDistDataRequest creates a request to invoke CleanDistData API
func CreateCleanDistDataRequest() (request *CleanDistDataRequest) {
	request = &CleanDistDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CleanDistData", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCleanDistDataResponse creates a response to parse from CleanDistData response
func CreateCleanDistDataResponse() (response *CleanDistDataResponse) {
	response = &CleanDistDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}