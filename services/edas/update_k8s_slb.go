package edas

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

// UpdateK8sSlb invokes the edas.UpdateK8sSlb API synchronously
func (client *Client) UpdateK8sSlb(request *UpdateK8sSlbRequest) (response *UpdateK8sSlbResponse, err error) {
	response = CreateUpdateK8sSlbResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateK8sSlbWithChan invokes the edas.UpdateK8sSlb API asynchronously
func (client *Client) UpdateK8sSlbWithChan(request *UpdateK8sSlbRequest) (<-chan *UpdateK8sSlbResponse, <-chan error) {
	responseChan := make(chan *UpdateK8sSlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateK8sSlb(request)
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

// UpdateK8sSlbWithCallback invokes the edas.UpdateK8sSlb API asynchronously
func (client *Client) UpdateK8sSlbWithCallback(request *UpdateK8sSlbRequest, callback func(response *UpdateK8sSlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateK8sSlbResponse
		var err error
		defer close(result)
		response, err = client.UpdateK8sSlb(request)
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

// UpdateK8sSlbRequest is the request struct for api UpdateK8sSlb
type UpdateK8sSlbRequest struct {
	*requests.RoaRequest
	Scheduler            string           `position:"Query" name:"Scheduler"`
	ServicePortInfos     string           `position:"Query" name:"ServicePortInfos"`
	SlbProtocol          string           `position:"Query" name:"SlbProtocol"`
	SlbName              string           `position:"Query" name:"SlbName"`
	Port                 string           `position:"Query" name:"Port"`
	AppId                string           `position:"Query" name:"AppId"`
	DisableForceOverride requests.Boolean `position:"Query" name:"DisableForceOverride"`
	Specification        string           `position:"Query" name:"Specification"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	Type                 string           `position:"Query" name:"Type"`
	TargetPort           string           `position:"Query" name:"TargetPort"`
}

// UpdateK8sSlbResponse is the response struct for api UpdateK8sSlb
type UpdateK8sSlbResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateK8sSlbRequest creates a request to invoke UpdateK8sSlb API
func CreateUpdateK8sSlbRequest() (request *UpdateK8sSlbRequest) {
	request = &UpdateK8sSlbRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateK8sSlb", "/pop/v5/k8s/acs/k8s_slb_binding", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateK8sSlbResponse creates a response to parse from UpdateK8sSlb response
func CreateUpdateK8sSlbResponse() (response *UpdateK8sSlbResponse) {
	response = &UpdateK8sSlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
