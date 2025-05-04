package main 

import (
    "cdk8s"
    "github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
    "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2"
    "constructs"
)

type MyChart struct {
    cdk8s.Chart
}

func NewMyChart(scope constructs.Construct, id string) cdk8s.Chart {
    chart := cdk8s.NewChart(scope, &id)

    // Deployment
    deployment := cdk8splus25.NewDeployment(chart, jsii.String("WebAppDeployment"), &cdk8splus25.DeploymentProps{
        Containers: &[]*cdk8splus25.ContainerProps{
            {
                Image: jsii.String("nginx:latest"),
                Port:  jsii.Number(80),
            },
        },
    })

    // Service
    service := cdk8splus25.NewService(chart, jsii.String("WebAppService"), &cdk8splus25.ServiceProps{
        Ports: &[]*cdk8splus25.ServicePort{
            {
                Port:       jsii.Number(80),
                TargetPort: jsii.Number(80),
            },
        },
        Metadata: &cdk8s.ApiObjectMetadata{
            Labels: &map[string]*string{
                "app": jsii.String("web-app"),
            },
        },
    })

    service.AddSelectorLabel("app", "web-app")
    service.ExposeDeployment(deployment)

    // Ingress
    ingress := cdk8splus25.NewIngress(chart, jsii.String("WebAppIngress"), &cdk8splus25.IngressProps{})
    ingress.AddHostDefaultBackend("/", deployment)

    return chart
}

func main() {
    app := cdk8s.NewApp(nil)
    NewMyChart(app, "go-web-app")
    app.Synth()
}
