package ltl

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

// SetDataWithSignature invokes the ltl.SetDataWithSignature API synchronously
func (client *Client) SetDataWithSignature(request *SetDataWithSignatureRequest) (response *SetDataWithSignatureResponse, err error) {
	response = CreateSetDataWithSignatureResponse()
	err = client.DoAction(request, response)
	return
}

// SetDataWithSignatureWithChan invokes the ltl.SetDataWithSignature API asynchronously
func (client *Client) SetDataWithSignatureWithChan(request *SetDataWithSignatureRequest) (<-chan *SetDataWithSignatureResponse, <-chan error) {
	responseChan := make(chan *SetDataWithSignatureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDataWithSignature(request)
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

// SetDataWithSignatureWithCallback invokes the ltl.SetDataWithSignature API asynchronously
func (client *Client) SetDataWithSignatureWithCallback(request *SetDataWithSignatureRequest, callback func(response *SetDataWithSignatureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDataWithSignatureResponse
		var err error
		defer close(result)
		response, err = client.SetDataWithSignature(request)
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

// SetDataWithSignatureRequest is the request struct for api SetDataWithSignature
type SetDataWithSignatureRequest struct {
	*requests.RpcRequest
	IotSignature         string `position:"Query" name:"IotSignature"`
	IotAuthType          string `position:"Query" name:"IotAuthType"`
	IotIdSource          string `position:"Query" name:"IotIdSource"`
	ApiVersion           string `position:"Query" name:"ApiVersion"`
	ProductKey           string `position:"Query" name:"ProductKey"`
	IotId                string `position:"Query" name:"IotId"`
	IotDataDigest        string `position:"Query" name:"IotDataDigest"`
	IotIdServiceProvider string `position:"Query" name:"IotIdServiceProvider"`
	Value                string `position:"Query" name:"Value"`
	Key                  string `position:"Query" name:"Key"`
}

// SetDataWithSignatureResponse is the response struct for api SetDataWithSignature
type SetDataWithSignatureResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSetDataWithSignatureRequest creates a request to invoke SetDataWithSignature API
func CreateSetDataWithSignatureRequest() (request *SetDataWithSignatureRequest) {
	request = &SetDataWithSignatureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "SetDataWithSignature", "", "")
	request.Method = requests.POST
	return
}

// CreateSetDataWithSignatureResponse creates a response to parse from SetDataWithSignature response
func CreateSetDataWithSignatureResponse() (response *SetDataWithSignatureResponse) {
	response = &SetDataWithSignatureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
