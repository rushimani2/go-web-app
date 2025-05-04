package main 

import (
  "cdk8s"
  "constructs"
  "github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
  k8s "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2"
)

type WebAppChart struct {
  cdk8s.Chart
}

func NewWebAppChart(scope constructs.Construct, id string) cdk8s.Chart {
  chart := cdk8s.NewChart(scope, &id)

  label := map[string]*string{"app": cdk8s.String("go-web-app")}

  deployment := k8s.NewDeployment(chart, cdk8s.String("go-web-deployment"), &k8s.DeploymentProps{
    Metadata: &cdk8s.ApiObjectMetadata{
      Name: cdk8s.String("go-web-deployment"),
    },
    Replicas: cdk8s.Float64(2),
    Selector: &cdk8s.LabelSelector{
      MatchLabels: &label,
    },
    PodMetadata: &cdk8s.ApiObjectMetadata{
      Labels: &label,
    },
    Containers: &[]*k8s.Container{
      {
        Image: cdk8s.String("nginx:latest"),
        Name:  cdk8s.String("nginx"),
        Ports: &[]*k8s.ContainerPort{
          {ContainerPort: cdk8s.Float64(80)},
        },
      },
    },
  })

  _ = deployment

  k8s.NewService(chart, cdk8s.String("go-web-service"), &k8s.ServiceProps{
    Metadata: &cdk8s.ApiObjectMetadata{
      Name: cdk8s.String("go-web-service"),
    },
    Selector: &label,
    Ports: &[]*k8s.ServicePort{
      {
        Port:       cdk8s.Float64(80),
        TargetPort: cdk8s.Float64(80),
      },
    },
  })

  return chart
}

func main() {
  app := cdk8s.NewApp(nil)
  NewWebAppChart(app, "go-web-app")
  app.Synth()
}
