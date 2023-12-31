package cloudapi

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

// DeleteApi invokes the cloudapi.DeleteApi API synchronously
func (client *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
	response = CreateDeleteApiResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApiWithChan invokes the cloudapi.DeleteApi API asynchronously
func (client *Client) DeleteApiWithChan(request *DeleteApiRequest) (<-chan *DeleteApiResponse, <-chan error) {
	responseChan := make(chan *DeleteApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApi(request)
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

// DeleteApiWithCallback invokes the cloudapi.DeleteApi API asynchronously
func (client *Client) DeleteApiWithCallback(request *DeleteApiRequest, callback func(response *DeleteApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApiResponse
		var err error
		defer close(result)
		response, err = client.DeleteApi(request)
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

// DeleteApiRequest is the request struct for api DeleteApi
type DeleteApiRequest struct {
	*requests.RpcRequest
	GroupId       string `position:"Query" name:"GroupId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DeleteApiResponse is the response struct for api DeleteApi
type DeleteApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteApiRequest creates a request to invoke DeleteApi API
func CreateDeleteApiRequest() (request *DeleteApiRequest) {
	request = &DeleteApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApi", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteApiResponse creates a response to parse from DeleteApi response
func CreateDeleteApiResponse() (response *DeleteApiResponse) {
	response = &DeleteApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
