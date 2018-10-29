package main

import (
	"context"
	"runtime"
	"time"

	stub "github.com/Percona-Lab/percona-server-mongodb-operator/pkg/stub"
	version "github.com/Percona-Lab/percona-server-mongodb-operator/version"
	opSdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	k8sutil "github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	mongodbOT "github.com/percona/mongodb-orchestration-tools"

	sdk "github.com/Percona-Lab/percona-server-mongodb-operator/pkg/sdk"

	"github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	GitCommit string
	GitBranch string
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
	logrus.Infof("perconalab/percona-server-mongodb-operator Version: %v, git commit: %s (branch: %s)", version.Version, GitCommit, GitBranch)
	logrus.Infof("percona/mongodb-orchestration-tools Version: %v", mongodbOT.Version)
}

func main() {
	printVersion()
	opSdk.ExposeMetricsPort()

	resource := "psmdb.percona.com/v1alpha1"
	kind := "PerconaServerMongoDB"
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logrus.Fatalf("failed to get watch namespace: %v", err)
	}

	resyncPeriod := time.Duration(5) * time.Second
	logrus.Infof("Watching %s, %s, %s, %s", resource, kind, namespace, resyncPeriod)
	opSdk.Watch(resource, kind, namespace, resyncPeriod)
	opSdk.Handle(stub.NewHandler(sdk.NewClient()))
	opSdk.Run(context.TODO())
}
