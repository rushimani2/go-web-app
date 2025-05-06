import { App, Chart } from 'cdk8s';
import { Construct } from 'constructs';
import * as k8s from 'cdk8s-plus-25';

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    const label = { app: 'web' };

    const deployment = new k8s.Deployment(this, 'web-deployment', {
      metadata: { name: 'web' }
    });

    deployment.addContainer({
      image: 'nginx',
      port: 80
    });

    // Add label selector for pods
    deployment.select.addMatchLabel('app', 'web');

    // Add the label to the pod template metadata so it matches the selector
    deployment.podMetadata.addLabel('app', 'web');

    new k8s.Service(this, 'web-service', {
      metadata: { name: 'web' },
      type: k8s.ServiceType.CLUSTER_IP,
      ports: [{ port: 80, targetPort: 80 }],
      selector: deployment.select.toJson()  // Corrected: use .toJson()
    });
  }
}

const app = new App();
new MyChart(app, 'my-chart');
app.synth();
