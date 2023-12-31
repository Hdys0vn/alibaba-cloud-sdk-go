package vs

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

// AddVsPullStreamInfoConfig invokes the vs.AddVsPullStreamInfoConfig API synchronously
func (client *Client) AddVsPullStreamInfoConfig(request *AddVsPullStreamInfoConfigRequest) (response *AddVsPullStreamInfoConfigResponse, err error) {
	response = CreateAddVsPullStreamInfoConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddVsPullStreamInfoConfigWithChan invokes the vs.AddVsPullStreamInfoConfig API asynchronously
func (client *Client) AddVsPullStreamInfoConfigWithChan(request *AddVsPullStreamInfoConfigRequest) (<-chan *AddVsPullStreamInfoConfigResponse, <-chan error) {
	responseChan := make(chan *AddVsPullStreamInfoConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddVsPullStreamInfoConfig(request)
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

// AddVsPullStreamInfoConfigWithCallback invokes the vs.AddVsPullStreamInfoConfig API asynchronously
func (client *Client) AddVsPullStreamInfoConfigWithCallback(request *AddVsPullStreamInfoConfigRequest, callback func(response *AddVsPullStreamInfoConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddVsPullStreamInfoConfigResponse
		var err error
		defer close(result)
		response, err = client.AddVsPullStreamInfoConfig(request)
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

// AddVsPullStreamInfoConfigRequest is the request struct for api AddVsPullStreamInfoConfig
type AddVsPullStreamInfoConfigRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	Always     string           `position:"Query" name:"Always"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	SourceUrl  string           `position:"Query" name:"SourceUrl"`
}

// AddVsPullStreamInfoConfigResponse is the response struct for api AddVsPullStreamInfoConfig
type AddVsPullStreamInfoConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddVsPullStreamInfoConfigRequest creates a request to invoke AddVsPullStreamInfoConfig API
func CreateAddVsPullStreamInfoConfigRequest() (request *AddVsPullStreamInfoConfigRequest) {
	request = &AddVsPullStreamInfoConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "AddVsPullStreamInfoConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateAddVsPullStreamInfoConfigResponse creates a response to parse from AddVsPullStreamInfoConfig response
func CreateAddVsPullStreamInfoConfigResponse() (response *AddVsPullStreamInfoConfigResponse) {
	response = &AddVsPullStreamInfoConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
