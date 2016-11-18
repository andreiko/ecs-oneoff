package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
)

func main() {
	taskDefinition := kingpin.Flag("taskdef", "Family and/or revision of the task definition to run").Required().String()
	cluster := kingpin.Flag("cluster", "Cluster on which to run your task").Default("default").String()
	count := kingpin.Flag("count", "The number of instantiations of each task to place on your cluster").Default("1").Int64()
	overrides := kingpin.Arg("override", "name of one or more override files in JSON format").Strings()

	kingpin.Parse()

	ecsClient := ecs.New(session.New())
	var inputs []*ecs.RunTaskInput

	if *count < 1 || *count > 10 {
		kingpin.Fatalf("count can only take values between 1 and 10")
	}

	if len(*overrides) > 0 {
		for _, filename := range *overrides {
			contents, err := ioutil.ReadFile(filename)
			if err != nil {
				kingpin.Fatalf("could not read %s: %s", filename, err.Error())
			}

			taskOverrides := new(ecs.TaskOverride)
			if err := json.Unmarshal(contents, taskOverrides); err != nil {
				kingpin.Fatalf("could not parse json from %s: %s", filename, err.Error())
			}

			inputs = append(inputs, &ecs.RunTaskInput{
				Cluster:        cluster,
				Count:          count,
				TaskDefinition: taskDefinition,
				Overrides:      taskOverrides,
			})
		}
	} else {
		// no overrides
		inputs = append(inputs, &ecs.RunTaskInput{
			Cluster:        cluster,
			Count:          count,
			TaskDefinition: taskDefinition,
		})
	}

	for _, input := range inputs {
		output, err := ecsClient.RunTask(input)
		if err != nil {
			kingpin.Errorf(err.Error())
			continue
		}

		for _, task := range output.Tasks {
			fmt.Println(*task.TaskArn)
		}
	}
}
