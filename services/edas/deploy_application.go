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

// DeployApplication invokes the edas.DeployApplication API synchronously
func (client *Client) DeployApplication(request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
	response = CreateDeployApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// DeployApplicationWithChan invokes the edas.DeployApplication API asynchronously
func (client *Client) DeployApplicationWithChan(request *DeployApplicationRequest) (<-chan *DeployApplicationResponse, <-chan error) {
	responseChan := make(chan *DeployApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployApplication(request)
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

// DeployApplicationWithCallback invokes the edas.DeployApplication API asynchronously
func (client *Client) DeployApplicationWithCallback(request *DeployApplicationRequest, callback func(response *DeployApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployApplicationResponse
		var err error
		defer close(result)
		response, err = client.DeployApplication(request)
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

// DeployApplicationRequest is the request struct for api DeployApplication
type DeployApplicationRequest struct {
	*requests.RoaRequest
	BuildPackId            requests.Integer `position:"Query" name:"BuildPackId"`
	ComponentIds           string           `position:"Query" name:"ComponentIds"`
	GroupId                string           `position:"Query" name:"GroupId"`
	BatchWaitTime          requests.Integer `position:"Query" name:"BatchWaitTime"`
	ReleaseType            requests.Integer `position:"Query" name:"ReleaseType"`
	Batch                  requests.Integer `position:"Query" name:"Batch"`
	AppEnv                 string           `position:"Query" name:"AppEnv"`
	PackageVersion         string           `position:"Query" name:"PackageVersion"`
	Gray                   requests.Boolean `position:"Query" name:"Gray"`
	AppId                  string           `position:"Query" name:"AppId"`
	ImageUrl               string           `position:"Query" name:"ImageUrl"`
	WarUrl                 string           `position:"Query" name:"WarUrl"`
	TrafficControlStrategy string           `position:"Query" name:"TrafficControlStrategy"`
	Desc                   string           `position:"Query" name:"Desc"`
	DeployType             string           `position:"Query" name:"DeployType"`
}

// DeployApplicationResponse is the response struct for api DeployApplication
type DeployApplicationResponse struct {
	*responses.BaseResponse
	Code          int    `json:"Code" xml:"Code"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeployApplicationRequest creates a request to invoke DeployApplication API
func CreateDeployApplicationRequest() (request *DeployApplicationRequest) {
	request = &DeployApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeployApplication", "/pop/v5/changeorder/co_deploy", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeployApplicationResponse creates a response to parse from DeployApplication response
func CreateDeployApplicationResponse() (response *DeployApplicationResponse) {
	response = &DeployApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
