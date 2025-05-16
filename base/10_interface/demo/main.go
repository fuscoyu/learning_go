package main

import "fmt"

// 1. 定义一个通用的 Client 接口
type Client interface {
	DoRequest(endpoint string, data interface{}) (interface{}, error)
}

// 2. 定义多个实现结构体
type HttpClient struct {
	BaseURL string
	Token   string
}

func (c *HttpClient) DoRequest(endpoint string, data interface{}) (interface{}, error) {
	// 这是一个示例，假设 HTTP 请求
	fmt.Printf("HttpClient 请求 %s，BaseURL: %s，Token: %s\n", endpoint, c.BaseURL, c.Token)
	return "HTTP 响应", nil
}

type GRPCClient struct {
	ServerAddress string
	Secure        bool
}

func (c *GRPCClient) DoRequest(endpoint string, data interface{}) (interface{}, error) {
	// 这是一个示例，假设 gRPC 请求
	fmt.Printf("GRPCClient 请求 %s，ServerAddress: %s，Secure: %v\n", endpoint, c.ServerAddress, c.Secure)
	return "gRPC 响应", nil
}

type CustomClient struct {
	APIKey string
}

func (c *CustomClient) DoRequest(endpoint string, data interface{}) (interface{}, error) {
	// 这是一个示例，假设自定义请求
	fmt.Printf("CustomClient 请求 %s，APIKey: %s\n", endpoint, c.APIKey)
	return "自定义响应", nil
}

// 3. 统一的请求方法
func MakeRequest(client Client, endpoint string, data interface{}) {
	response, err := client.DoRequest(endpoint, data)
	if err != nil {
		fmt.Println("请求失败：", err)
	} else {
		fmt.Println("请求成功，响应：", response)
	}
}

func main() {
	httpClient := &HttpClient{BaseURL: "https://api.example.com", Token: "abc123"}
	grpcClient := &GRPCClient{ServerAddress: "localhost:50051", Secure: true}
	customClient := &CustomClient{APIKey: "my-secret-api-key"}

	// 统一的请求调用
	MakeRequest(httpClient, "/v1/users", nil)
	MakeRequest(grpcClient, "/v1/users", nil)
	MakeRequest(customClient, "/v1/users", nil)
}
