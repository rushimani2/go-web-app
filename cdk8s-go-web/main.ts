import { App, Chart } from 'cdk8s';
import { Construct } from 'constructs';
import * as k8s from 'cdk8s-plus-25';

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    // Create Deployment
    const deployment = new k8s.Deployment(this, 'web-deployment', {
      metadata: { name: 'web' },
      containers: [{ image: 'nginx' }],
    });

    // Create Service
    new k8s.Service(this, 'web-service', {
      metadata: { name: 'web-service' },
      selector: { app: 'web' }, // Match Deployment selector
      ports: [{ port: 80 }], // Expose port 80
    });

    // Create Ingress
    new k8s.Ingress(this, 'web-ingress', {
      metadata: { name: 'web-ingress' },
      rules: [
        {
          host: 'web.local', // Change this according to your requirements
          http: {
            paths: [
              {
                path: '/',
                backend: {
                  service: {
                    name: 'web-service',
                    port: 80, // Correcting the port definition
                  },
                },
              },
            ],
          },
        },
      ],
    });
  }
}

const app = new App();
new MyChart(app, 'my-chart');
app.synth();
