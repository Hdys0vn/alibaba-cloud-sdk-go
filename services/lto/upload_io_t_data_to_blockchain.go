package lto

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

// UploadIoTDataToBlockchain invokes the lto.UploadIoTDataToBlockchain API synchronously
func (client *Client) UploadIoTDataToBlockchain(request *UploadIoTDataToBlockchainRequest) (response *UploadIoTDataToBlockchainResponse, err error) {
	response = CreateUploadIoTDataToBlockchainResponse()
	err = client.DoAction(request, response)
	return
}

// UploadIoTDataToBlockchainWithChan invokes the lto.UploadIoTDataToBlockchain API asynchronously
func (client *Client) UploadIoTDataToBlockchainWithChan(request *UploadIoTDataToBlockchainRequest) (<-chan *UploadIoTDataToBlockchainResponse, <-chan error) {
	responseChan := make(chan *UploadIoTDataToBlockchainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadIoTDataToBlockchain(request)
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

// UploadIoTDataToBlockchainWithCallback invokes the lto.UploadIoTDataToBlockchain API asynchronously
func (client *Client) UploadIoTDataToBlockchainWithCallback(request *UploadIoTDataToBlockchainRequest, callback func(response *UploadIoTDataToBlockchainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadIoTDataToBlockchainResponse
		var err error
		defer close(result)
		response, err = client.UploadIoTDataToBlockchain(request)
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

// UploadIoTDataToBlockchainRequest is the request struct for api UploadIoTDataToBlockchain
type UploadIoTDataToBlockchainRequest struct {
	*requests.RpcRequest
	IotIdSource          string `position:"Query" name:"IotIdSource"`
	IotDataToken         string `position:"Query" name:"IotDataToken"`
	PrivacyData          string `position:"Query" name:"PrivacyData"`
	IotId                string `position:"Query" name:"IotId"`
	IotDataDigest        string `position:"Query" name:"IotDataDigest"`
	IotDataDID           string `position:"Query" name:"IotDataDID"`
	PlainData            string `position:"Query" name:"PlainData"`
	IotAuthType          string `position:"Query" name:"IotAuthType"`
	IotIdServiceProvider string `position:"Query" name:"IotIdServiceProvider"`
}

// UploadIoTDataToBlockchainResponse is the response struct for api UploadIoTDataToBlockchain
type UploadIoTDataToBlockchainResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUploadIoTDataToBlockchainRequest creates a request to invoke UploadIoTDataToBlockchain API
func CreateUploadIoTDataToBlockchainRequest() (request *UploadIoTDataToBlockchainRequest) {
	request = &UploadIoTDataToBlockchainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "UploadIoTDataToBlockchain", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadIoTDataToBlockchainResponse creates a response to parse from UploadIoTDataToBlockchain response
func CreateUploadIoTDataToBlockchainResponse() (response *UploadIoTDataToBlockchainResponse) {
	response = &UploadIoTDataToBlockchainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
