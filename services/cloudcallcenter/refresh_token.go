package cloudcallcenter

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

// RefreshToken invokes the cloudcallcenter.RefreshToken API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/refreshtoken.html
func (client *Client) RefreshToken(request *RefreshTokenRequest) (response *RefreshTokenResponse, err error) {
	response = CreateRefreshTokenResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshTokenWithChan invokes the cloudcallcenter.RefreshToken API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/refreshtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshTokenWithChan(request *RefreshTokenRequest) (<-chan *RefreshTokenResponse, <-chan error) {
	responseChan := make(chan *RefreshTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshToken(request)
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

// RefreshTokenWithCallback invokes the cloudcallcenter.RefreshToken API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/refreshtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshTokenWithCallback(request *RefreshTokenRequest, callback func(response *RefreshTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshTokenResponse
		var err error
		defer close(result)
		response, err = client.RefreshToken(request)
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

// RefreshTokenRequest is the request struct for api RefreshToken
type RefreshTokenRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// RefreshTokenResponse is the response struct for api RefreshToken
type RefreshTokenResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Token          Token  `json:"Token" xml:"Token"`
}

// CreateRefreshTokenRequest creates a request to invoke RefreshToken API
func CreateRefreshTokenRequest() (request *RefreshTokenRequest) {
	request = &RefreshTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "RefreshToken", "", "")
	request.Method = requests.POST
	return
}

// CreateRefreshTokenResponse creates a response to parse from RefreshToken response
func CreateRefreshTokenResponse() (response *RefreshTokenResponse) {
	response = &RefreshTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
