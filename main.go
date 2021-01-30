package main

import (
	"github.com/pulumi/pulumi-packet/sdk/v3/go/packet"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		c := config.New(ctx, "")
		// Create a Packet resource (Project)

		project := c.Require("projectId")

		nginx, err := packet.NewDevice(ctx, "nginx-server", &packet.DeviceArgs{
			ProjectId:       pulumi.String(project),
			Hostname:        pulumi.String("ngnix-server"),
			Facilities:      pulumi.StringArray{pulumi.String("am6")},
			BillingCycle:    pulumi.String("hourly"),
			OperatingSystem: pulumi.String("ubuntu_18_04"),
			Plan:            pulumi.String("c3.small.x86"),
			UserData:        pulumi.String("#!/usr/bin/env bash\necho 'Hello'"),
		})
		if err != nil {
			return err
		}

		ctx.Export("nginxIp", nginx.AccessPrivateIpv4)
		return nil
	})
}
