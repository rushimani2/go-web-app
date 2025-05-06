import { App, Chart } from "cdk8s";
import { Construct } from "constructs";
import * as k8s from "cdk8s-plus-25";

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    // Define the label selector
    const appLabel = { app: "go-web-app" };

    // Create the Deployment
    new k8s.Deployment(this, "go-web-app-deployment", {
      metadata: {
        name: "go-web-app",
        labels: appLabel,
      },
      spec: {
        replicas: 1,
        selector: k8s.LabelSelector.all(appLabel),  // Correct selector for matching labels
        template: {
          metadata: {
            labels: appLabel,
          },
          spec: {
            containers: [
              {
                name: "go-web-app",
                image: "rushibindu/go-web-app:{{ .Values.image.tag }}",
                ports: [{ containerPort: 8080 }],
              },
            ],
          },
        },
      },
    });

    // Create the Service
    new k8s.Service(this, "go-web-app-service", {
      metadata: {
        name: "go-web-app",
        labels: appLabel,
      },
      spec: {
        ports: [
          {
            port: 80,
            targetPort: 8080,
            protocol: k8s.Protocol.TCP,
          },
        ],
        selector: appLabel,  // Correctly apply the label selector for the service
        type: k8s.ServiceType.LOAD_BALANCER,
      },
    });
  }
}

const app = new App();
new MyChart(app, "go-web-app-chart");
app.synth();
