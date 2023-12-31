package dyplsapi

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

// BindBatchAxg invokes the dyplsapi.BindBatchAxg API synchronously
func (client *Client) BindBatchAxg(request *BindBatchAxgRequest) (response *BindBatchAxgResponse, err error) {
	response = CreateBindBatchAxgResponse()
	err = client.DoAction(request, response)
	return
}

// BindBatchAxgWithChan invokes the dyplsapi.BindBatchAxg API asynchronously
func (client *Client) BindBatchAxgWithChan(request *BindBatchAxgRequest) (<-chan *BindBatchAxgResponse, <-chan error) {
	responseChan := make(chan *BindBatchAxgResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindBatchAxg(request)
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

// BindBatchAxgWithCallback invokes the dyplsapi.BindBatchAxg API asynchronously
func (client *Client) BindBatchAxgWithCallback(request *BindBatchAxgRequest, callback func(response *BindBatchAxgResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindBatchAxgResponse
		var err error
		defer close(result)
		response, err = client.BindBatchAxg(request)
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

// BindBatchAxgRequest is the request struct for api BindBatchAxg
type BindBatchAxgRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer           `position:"Query" name:"ResourceOwnerId"`
	AxgBindList          *[]BindBatchAxgAxgBindList `position:"Query" name:"AxgBindList"  type:"Json"`
	ResourceOwnerAccount string                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer           `position:"Query" name:"OwnerId"`
	PoolKey              string                     `position:"Query" name:"PoolKey"`
}

// BindBatchAxgAxgBindList is a repeated param struct in BindBatchAxgRequest
type BindBatchAxgAxgBindList struct {
	PhoneNoB           string `name:"PhoneNoB"`
	PhoneNoA           string `name:"PhoneNoA"`
	ExpectCity         string `name:"ExpectCity"`
	GroupId            string `name:"GroupId"`
	CallDisplayType    string `name:"CallDisplayType"`
	OutOrderId         string `name:"OutOrderId"`
	PhoneNoX           string `name:"PhoneNoX"`
	IsRecordingEnabled string `name:"IsRecordingEnabled"`
	OutId              string `name:"OutId"`
	Expiration         string `name:"Expiration"`
	RingConfig         string `name:"RingConfig"`
	ASRStatus          string `name:"ASRStatus"`
	ASRModelId         string `name:"ASRModelId"`
	CallRestrict       string `name:"CallRestrict"`
}

// BindBatchAxgResponse is the response struct for api BindBatchAxg
type BindBatchAxgResponse struct {
	*responses.BaseResponse
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	SecretBindList SecretBindList `json:"SecretBindList" xml:"SecretBindList"`
}

// CreateBindBatchAxgRequest creates a request to invoke BindBatchAxg API
func CreateBindBatchAxgRequest() (request *BindBatchAxgRequest) {
	request = &BindBatchAxgRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "BindBatchAxg", "", "")
	request.Method = requests.POST
	return
}

// CreateBindBatchAxgResponse creates a response to parse from BindBatchAxg response
func CreateBindBatchAxgResponse() (response *BindBatchAxgResponse) {
	response = &BindBatchAxgResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
