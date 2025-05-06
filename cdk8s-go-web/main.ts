import { App, Chart } from 'cdk8s';
import { Construct } from 'constructs';
import * as k8s from 'cdk8s-plus-25';

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    const appLabel = { app: 'go-web-app' };

    // Create Deployment
    const deployment = new k8s.Deployment(this, 'go-web-app-deployment', {
      metadata: {
        name: 'go-web-app',
        labels: appLabel,
      },
      replicas: 1,
      selector: { matchLabels: appLabel },
    });

    // Add container to the Deployment
    deployment.addContainer({
      name: 'go-web-app',
      image: 'rushibindu/go-web-app:{{ .Values.image.tag }}',
      ports: [{ containerPort: 8080 }],
    });

    // Create Service
    new k8s.Service(this, 'go-web-app-service', {
      metadata: {
        name: 'go-web-app',
        labels: appLabel,
      },
      ports: [
        {
          port: 80,
          targetPort: 8080,
          protocol: k8s.Protocol.TCP,
        },
      ],
      selector: appLabel,
      type: k8s.ServiceType.LOAD_BALANCER,
    });
  }
}

const app = new App();
new MyChart(app, 'go-web-app-chart');
app.synth();
