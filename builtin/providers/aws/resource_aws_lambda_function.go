package aws

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/awslabs/aws-sdk-go/aws"
	_ "github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/lambda"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsLambdaFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsLambdaFunctionCreate,
		Read:   resourceAwsLambdaFunctionRead,
		Update: resourceAwsLambdaFunctionUpdate,
		Delete: resourceAwsLambdaFunctionDelete,

		Schema: map[string]*schema.Schema{
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"function_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"handler": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"runtime": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
		},
	}
}

// readZip reads a zipfile and returns the zip data as a base64-encoded byte
// array so you can upload it to the Lambda service.
func readZip(filename string) ([]byte, error) {

	data, err := ioutil.ReadFile(filename)
	b64size = ceil(len(data)/3) * 4
	target := make([]byte, b64size)
	if err != nil {
		return target, err
	}
	base64.StdEncoding.Encode(target, data)
	return target, nil
}

// resourceAwsLambdaFunction maps to:
// CreateFunction in the API / SDK
func resourceAwsLambdaFunctionCreate(d *schema.ResourceData, meta interface{}) error {

	svc := lambda.New(nil)

	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			ZipFile: []byte(d.Get("code").(string)),
		},
		Description:  aws.String(d.Get("description").(string)),
		FunctionName: aws.String(d.Get("function_name").(string)),
		Handler:      aws.String(d.Get("handler").(string)),
		MemorySize:   aws.Long(d.Get("memory_size").(int64)),
		Role:         aws.String(d.Get("role").(string)),
		Runtime:      aws.String(d.Get("runtime").(string)),
		Timeout:      aws.Long(d.Get("timeout").(int64)),
	}

	resp, err := svc.CreateFunction(params)

	functionName := resp.FunctionName
	fmt.Println(functionName)
	fmt.Println(resp)

	// Do something with resp

	return err
}

// resourceAwsLambdaFunctionRead maps to:
// GetFunction in the API / SDK
func resourceAwsLambdaFunctionRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceAwsLambdaFunction maps to:
// DeleteFunction in the API / SDK
func resourceAwsLambdaFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceAwsLambdaFunctionUpdate maps to:
// UpdateFunctionCode in the API / SDK
func resourceAwsLambdaFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
