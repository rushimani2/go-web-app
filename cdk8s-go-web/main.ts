import { App, Chart } from 'cdk8s';
import { Construct } from 'constructs';
import * as k8s from 'cdk8s-plus-25';

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    // Create a Deployment resource
    new k8s.Deployment(this, 'web-deployment', {
      metadata: { name: 'web' },
      containers: [
        {
          image: 'nginx',  // The container image for the Deployment
          name: 'nginx-container',
        }
      ]
    });
  }
}

const app = new App();
new MyChart(app, 'my-chart');
app.synth();
