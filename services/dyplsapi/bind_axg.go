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

// BindAxg invokes the dyplsapi.BindAxg API synchronously
func (client *Client) BindAxg(request *BindAxgRequest) (response *BindAxgResponse, err error) {
	response = CreateBindAxgResponse()
	err = client.DoAction(request, response)
	return
}

// BindAxgWithChan invokes the dyplsapi.BindAxg API asynchronously
func (client *Client) BindAxgWithChan(request *BindAxgRequest) (<-chan *BindAxgResponse, <-chan error) {
	responseChan := make(chan *BindAxgResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindAxg(request)
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

// BindAxgWithCallback invokes the dyplsapi.BindAxg API asynchronously
func (client *Client) BindAxgWithCallback(request *BindAxgRequest, callback func(response *BindAxgResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindAxgResponse
		var err error
		defer close(result)
		response, err = client.BindAxg(request)
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

// BindAxgRequest is the request struct for api BindAxg
type BindAxgRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CallDisplayType      requests.Integer `position:"Query" name:"CallDisplayType"`
	PhoneNoX             string           `position:"Query" name:"PhoneNoX"`
	RingConfig           string           `position:"Query" name:"RingConfig"`
	ASRStatus            requests.Boolean `position:"Query" name:"ASRStatus"`
	PhoneNoB             string           `position:"Query" name:"PhoneNoB"`
	PhoneNoA             string           `position:"Query" name:"PhoneNoA"`
	ExpectCity           string           `position:"Query" name:"ExpectCity"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	GroupId              string           `position:"Query" name:"GroupId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OutOrderId           string           `position:"Query" name:"OutOrderId"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
	Expiration           string           `position:"Query" name:"Expiration"`
	IsRecordingEnabled   requests.Boolean `position:"Query" name:"IsRecordingEnabled"`
	OutId                string           `position:"Query" name:"OutId"`
	ASRModelId           string           `position:"Query" name:"ASRModelId"`
	CallRestrict         string           `position:"Query" name:"CallRestrict"`
}

// BindAxgResponse is the response struct for api BindAxg
type BindAxgResponse struct {
	*responses.BaseResponse
	Code          string        `json:"Code" xml:"Code"`
	Message       string        `json:"Message" xml:"Message"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	SecretBindDTO SecretBindDTO `json:"SecretBindDTO" xml:"SecretBindDTO"`
}

// CreateBindAxgRequest creates a request to invoke BindAxg API
func CreateBindAxgRequest() (request *BindAxgRequest) {
	request = &BindAxgRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "BindAxg", "", "")
	request.Method = requests.POST
	return
}

// CreateBindAxgResponse creates a response to parse from BindAxg response
func CreateBindAxgResponse() (response *BindAxgResponse) {
	response = &BindAxgResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
