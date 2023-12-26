package linkvisual

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

// UpdateInstanceInternetProtocol invokes the linkvisual.UpdateInstanceInternetProtocol API synchronously
func (client *Client) UpdateInstanceInternetProtocol(request *UpdateInstanceInternetProtocolRequest) (response *UpdateInstanceInternetProtocolResponse, err error) {
	response = CreateUpdateInstanceInternetProtocolResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateInstanceInternetProtocolWithChan invokes the linkvisual.UpdateInstanceInternetProtocol API asynchronously
func (client *Client) UpdateInstanceInternetProtocolWithChan(request *UpdateInstanceInternetProtocolRequest) (<-chan *UpdateInstanceInternetProtocolResponse, <-chan error) {
	responseChan := make(chan *UpdateInstanceInternetProtocolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateInstanceInternetProtocol(request)
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

// UpdateInstanceInternetProtocolWithCallback invokes the linkvisual.UpdateInstanceInternetProtocol API asynchronously
func (client *Client) UpdateInstanceInternetProtocolWithCallback(request *UpdateInstanceInternetProtocolRequest, callback func(response *UpdateInstanceInternetProtocolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateInstanceInternetProtocolResponse
		var err error
		defer close(result)
		response, err = client.UpdateInstanceInternetProtocol(request)
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

// UpdateInstanceInternetProtocolRequest is the request struct for api UpdateInstanceInternetProtocol
type UpdateInstanceInternetProtocolRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	IpVersion     string `position:"Query" name:"IpVersion"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// UpdateInstanceInternetProtocolResponse is the response struct for api UpdateInstanceInternetProtocol
type UpdateInstanceInternetProtocolResponse struct {
	*responses.BaseResponse
	Code         string                 `json:"Code" xml:"Code"`
	Data         map[string]interface{} `json:"Data" xml:"Data"`
	ErrorMessage string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	Success      bool                   `json:"Success" xml:"Success"`
}

// CreateUpdateInstanceInternetProtocolRequest creates a request to invoke UpdateInstanceInternetProtocol API
func CreateUpdateInstanceInternetProtocolRequest() (request *UpdateInstanceInternetProtocolRequest) {
	request = &UpdateInstanceInternetProtocolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "UpdateInstanceInternetProtocol", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateInstanceInternetProtocolResponse creates a response to parse from UpdateInstanceInternetProtocol response
func CreateUpdateInstanceInternetProtocolResponse() (response *UpdateInstanceInternetProtocolResponse) {
	response = &UpdateInstanceInternetProtocolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
