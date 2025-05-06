import { App, Chart } from 'cdk8s';
import { Construct } from 'constructs';
import * as k8s from 'cdk8s-plus-25';

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    // Create a Deployment
    new k8s.Deployment(this, 'web-deployment', {
      metadata: { name: 'web' },
      containers: [{ image: 'nginx' }],
    });

    // Create a Service
    new k8s.Service(this, 'web-service', {
      metadata: { name: 'web-service' },
      spec: {
        ports: [{ port: 80 }],
        selector: { app: 'web' }, // match the Deployment
      },
    });

    // Create an Ingress
    new k8s.Ingress(this, 'web-ingress', {
      metadata: { name: 'web-ingress' },
      spec: {
        rules: [
          {
            host: 'web.local', // You should change this for your real use case
            http: {
              paths: [
                {
                  path: '/',
                  backend: {
                    service: {
                      name: 'web-service',
                      port: { number: 80 },
                    },
                  },
                },
              ],
            },
          },
        ],
      },
    });
  }
}

const app = new App();
new MyChart(app, 'my-chart');
app.synth();
