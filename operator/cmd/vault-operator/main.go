package main

import (
	"context"
	"runtime"

	stub "github.com/banzaicloud/bank-vaults/operator/pkg/stub"
	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logrus.Fatalf("Failed to get watch namespace: %v", err)
	}
	logrus.Infof("watching namespace: %v", namespace)
	sdk.Watch("vault.banzaicloud.com/v1alpha1", "Vault", namespace, 60)
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
