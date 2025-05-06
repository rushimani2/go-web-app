import { App, Chart } from 'cdk8s';
import { Construct } from 'constructs';
import * as k8s from 'cdk8s-plus-25';

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    const label = { app: 'web' };

    // Create a deployment with the label
    const deployment = new k8s.Deployment(this, 'web-deployment', {
      metadata: { name: 'web' },
    });

    // Add a container to the deployment
    deployment.addContainer({
      image: 'nginx',
      port: 80,
    });

    // Set label on the deployment's pod metadata to match the service selector
    deployment.podMetadata.addLabel('app', 'web');

    // Create a Kubernetes service that selects the pods using the 'app' label
    new k8s.Service(this, 'web-service', {
      metadata: { name: 'web' },
      type: k8s.ServiceType.CLUSTER_IP,
      ports: [{ port: 80, targetPort: 80 }],
      selector: label, // Use the same label for the service selector
    });
  }
}

const app = new App();
new MyChart(app, 'my-chart');
app.synth();
