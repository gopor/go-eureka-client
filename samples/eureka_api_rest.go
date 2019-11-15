package main

import (
    "github.com/gopor/go-eureka-client/eureka"
    "log"
)

func main() {
    config := eureka.GetDefaultEurekaClientConfig()
    config.UseDnsForFetchingServiceUrls = false
    config.Region = eureka.DEFAULT_REGION
    config.AvailabilityZones = map[string]string{
        eureka.DEFAULT_REGION: eureka.DEFAULT_ZONE,
    }
    config.ServiceUrl = map[string]string{
        eureka.DEFAULT_ZONE: "http://cmd:cmd123@192.168.1.15:8761/eureka,http://cmd:cmd123@192.168.1.15:8762/eureka",
    }

    c := eureka.DefaultClient.Config(config)
    api, err := c.Api()
    if err != nil {
        log.Fatalln("Failed to pick EurekaServerApi instance, err=", err.Error())
    }
    instances, err := api.QueryAllInstances()
    if err != nil {
        log.Fatalln("Failed to query all instances, err=", err.Error())
    }

    log.Println("all instances: ", instances)
}
