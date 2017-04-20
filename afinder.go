package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

// This example will list instances with a filter
//
// Usage:
// filter_ec2_by_tag <name_filter>
func main() {
	sess := session.Must(session.NewSession())

	nameFilter := os.Args[1]
	awsRegion := "us-east-1"
	svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	fmt.Printf("listing instances with tag %v in: %v\n", nameFilter, awsRegion)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", nameFilter, "*"}, "")),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", awsRegion, err.Error())
		log.Fatal(err.Error())
	}

	insid := *resp.Reservations[0].Instances[0].InstanceId
	privp := *resp.Reservations[0].Instances[0].PrivateIpAddress

	data := [][]string{
		[]string{insid, privp},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Instance ID", "PrivateIP"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	//fmt.Println(insid, privp)

}
