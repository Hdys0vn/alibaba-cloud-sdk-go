package linkwan

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

// UpdateGatewaySshCtrl invokes the linkwan.UpdateGatewaySshCtrl API synchronously
func (client *Client) UpdateGatewaySshCtrl(request *UpdateGatewaySshCtrlRequest) (response *UpdateGatewaySshCtrlResponse, err error) {
	response = CreateUpdateGatewaySshCtrlResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewaySshCtrlWithChan invokes the linkwan.UpdateGatewaySshCtrl API asynchronously
func (client *Client) UpdateGatewaySshCtrlWithChan(request *UpdateGatewaySshCtrlRequest) (<-chan *UpdateGatewaySshCtrlResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewaySshCtrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewaySshCtrl(request)
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

// UpdateGatewaySshCtrlWithCallback invokes the linkwan.UpdateGatewaySshCtrl API asynchronously
func (client *Client) UpdateGatewaySshCtrlWithCallback(request *UpdateGatewaySshCtrlRequest, callback func(response *UpdateGatewaySshCtrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewaySshCtrlResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewaySshCtrl(request)
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

// UpdateGatewaySshCtrlRequest is the request struct for api UpdateGatewaySshCtrl
type UpdateGatewaySshCtrlRequest struct {
	*requests.RpcRequest
	Enabled       requests.Boolean `position:"Query" name:"Enabled"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	GwEui         string           `position:"Query" name:"GwEui"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// UpdateGatewaySshCtrlResponse is the response struct for api UpdateGatewaySshCtrl
type UpdateGatewaySshCtrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateGatewaySshCtrlRequest creates a request to invoke UpdateGatewaySshCtrl API
func CreateUpdateGatewaySshCtrlRequest() (request *UpdateGatewaySshCtrlRequest) {
	request = &UpdateGatewaySshCtrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "UpdateGatewaySshCtrl", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewaySshCtrlResponse creates a response to parse from UpdateGatewaySshCtrl response
func CreateUpdateGatewaySshCtrlResponse() (response *UpdateGatewaySshCtrlResponse) {
	response = &UpdateGatewaySshCtrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}