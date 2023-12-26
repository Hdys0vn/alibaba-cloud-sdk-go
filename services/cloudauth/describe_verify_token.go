package cloudauth

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

// DescribeVerifyToken invokes the cloudauth.DescribeVerifyToken API synchronously
func (client *Client) DescribeVerifyToken(request *DescribeVerifyTokenRequest) (response *DescribeVerifyTokenResponse, err error) {
	response = CreateDescribeVerifyTokenResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVerifyTokenWithChan invokes the cloudauth.DescribeVerifyToken API asynchronously
func (client *Client) DescribeVerifyTokenWithChan(request *DescribeVerifyTokenRequest) (<-chan *DescribeVerifyTokenResponse, <-chan error) {
	responseChan := make(chan *DescribeVerifyTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVerifyToken(request)
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

// DescribeVerifyTokenWithCallback invokes the cloudauth.DescribeVerifyToken API asynchronously
func (client *Client) DescribeVerifyTokenWithCallback(request *DescribeVerifyTokenRequest, callback func(response *DescribeVerifyTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVerifyTokenResponse
		var err error
		defer close(result)
		response, err = client.DescribeVerifyToken(request)
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

// DescribeVerifyTokenRequest is the request struct for api DescribeVerifyToken
type DescribeVerifyTokenRequest struct {
	*requests.RpcRequest
	FaceRetainedImageUrl string           `position:"Query" name:"FaceRetainedImageUrl"`
	UserId               string           `position:"Query" name:"UserId"`
	CallbackSeed         string           `position:"Query" name:"CallbackSeed"`
	UserIp               string           `position:"Query" name:"UserIp"`
	IdCardBackImageUrl   string           `position:"Query" name:"IdCardBackImageUrl"`
	IdCardNumber         string           `position:"Query" name:"IdCardNumber"`
	IdCardFrontImageUrl  string           `position:"Query" name:"IdCardFrontImageUrl"`
	BizType              string           `position:"Query" name:"BizType"`
	PassedRedirectUrl    string           `position:"Query" name:"PassedRedirectUrl"`
	UserRegistTime       requests.Integer `position:"Query" name:"UserRegistTime"`
	BizId                string           `position:"Query" name:"BizId"`
	Name                 string           `position:"Query" name:"Name"`
	UserPhoneNumber      string           `position:"Query" name:"UserPhoneNumber"`
	CallbackUrl          string           `position:"Query" name:"CallbackUrl"`
	FailedRedirectUrl    string           `position:"Query" name:"FailedRedirectUrl"`
}

// DescribeVerifyTokenResponse is the response struct for api DescribeVerifyToken
type DescribeVerifyTokenResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	VerifyPageUrl  string         `json:"VerifyPageUrl" xml:"VerifyPageUrl"`
	VerifyToken    string         `json:"VerifyToken" xml:"VerifyToken"`
	OssUploadToken OssUploadToken `json:"OssUploadToken" xml:"OssUploadToken"`
}

// CreateDescribeVerifyTokenRequest creates a request to invoke DescribeVerifyToken API
func CreateDescribeVerifyTokenRequest() (request *DescribeVerifyTokenRequest) {
	request = &DescribeVerifyTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeVerifyToken", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVerifyTokenResponse creates a response to parse from DescribeVerifyToken response
func CreateDescribeVerifyTokenResponse() (response *DescribeVerifyTokenResponse) {
	response = &DescribeVerifyTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
