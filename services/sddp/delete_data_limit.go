package sddp

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

// DeleteDataLimit invokes the sddp.DeleteDataLimit API synchronously
func (client *Client) DeleteDataLimit(request *DeleteDataLimitRequest) (response *DeleteDataLimitResponse, err error) {
	response = CreateDeleteDataLimitResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDataLimitWithChan invokes the sddp.DeleteDataLimit API asynchronously
func (client *Client) DeleteDataLimitWithChan(request *DeleteDataLimitRequest) (<-chan *DeleteDataLimitResponse, <-chan error) {
	responseChan := make(chan *DeleteDataLimitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDataLimit(request)
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

// DeleteDataLimitWithCallback invokes the sddp.DeleteDataLimit API asynchronously
func (client *Client) DeleteDataLimitWithCallback(request *DeleteDataLimitRequest, callback func(response *DeleteDataLimitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDataLimitResponse
		var err error
		defer close(result)
		response, err = client.DeleteDataLimit(request)
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

// DeleteDataLimitRequest is the request struct for api DeleteDataLimit
type DeleteDataLimitRequest struct {
	*requests.RpcRequest
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DeleteDataLimitResponse is the response struct for api DeleteDataLimit
type DeleteDataLimitResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDataLimitRequest creates a request to invoke DeleteDataLimit API
func CreateDeleteDataLimitRequest() (request *DeleteDataLimitRequest) {
	request = &DeleteDataLimitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DeleteDataLimit", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDataLimitResponse creates a response to parse from DeleteDataLimit response
func CreateDeleteDataLimitResponse() (response *DeleteDataLimitResponse) {
	response = &DeleteDataLimitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}