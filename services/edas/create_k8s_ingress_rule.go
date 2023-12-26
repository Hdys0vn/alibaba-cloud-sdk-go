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

// CreateK8sIngressRule invokes the edas.CreateK8sIngressRule API synchronously
func (client *Client) CreateK8sIngressRule(request *CreateK8sIngressRuleRequest) (response *CreateK8sIngressRuleResponse, err error) {
	response = CreateCreateK8sIngressRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateK8sIngressRuleWithChan invokes the edas.CreateK8sIngressRule API asynchronously
func (client *Client) CreateK8sIngressRuleWithChan(request *CreateK8sIngressRuleRequest) (<-chan *CreateK8sIngressRuleResponse, <-chan error) {
	responseChan := make(chan *CreateK8sIngressRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateK8sIngressRule(request)
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

// CreateK8sIngressRuleWithCallback invokes the edas.CreateK8sIngressRule API asynchronously
func (client *Client) CreateK8sIngressRuleWithCallback(request *CreateK8sIngressRuleRequest, callback func(response *CreateK8sIngressRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateK8sIngressRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateK8sIngressRule(request)
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

// CreateK8sIngressRuleRequest is the request struct for api CreateK8sIngressRule
type CreateK8sIngressRuleRequest struct {
	*requests.RoaRequest
	Namespace   string `position:"Query" name:"Namespace"`
	Name        string `position:"Query" name:"Name"`
	IngressConf string `position:"Query" name:"IngressConf"`
	Annotations string `position:"Query" name:"Annotations"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	Labels      string `position:"Query" name:"Labels"`
}

// CreateK8sIngressRuleResponse is the response struct for api CreateK8sIngressRule
type CreateK8sIngressRuleResponse struct {
	*responses.BaseResponse
	Code    int    `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateCreateK8sIngressRuleRequest creates a request to invoke CreateK8sIngressRule API
func CreateCreateK8sIngressRuleRequest() (request *CreateK8sIngressRuleRequest) {
	request = &CreateK8sIngressRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "CreateK8sIngressRule", "/pop/v5/k8s/acs/k8s_ingress", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateK8sIngressRuleResponse creates a response to parse from CreateK8sIngressRule response
func CreateCreateK8sIngressRuleResponse() (response *CreateK8sIngressRuleResponse) {
	response = &CreateK8sIngressRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
