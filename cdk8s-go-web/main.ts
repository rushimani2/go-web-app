import { App, Chart } from 'cdk8s';
import { Construct } from 'constructs';
import * as k8s from 'cdk8s-plus-25';

class MyChart extends Chart {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    // Label used for Deployment and matched by the Service selector
    const label = { 'cdk8s.io/metadata.addr': 'my-chart-web-deployment-c887ba4c' };

    // Define the Deployment
    new k8s.Deployment(this, 'web-deployment', {
      metadata: { name: 'web' },
      replicas: 2,
      containers: [
        new k8s.Container({
          name: 'main',
          image: 'nginx',
          resources: {
            limits: {
              cpu: k8s.Cpu.millis(1500),
              memory: k8s.Size.mebibytes(2048),
            },
            requests: {
              cpu: k8s.Cpu.millis(1000),
              memory: k8s.Size.mebibytes(512),
            },
          },
          securityContext: {
            allowPrivilegeEscalation: false,
            privileged: false,
            readOnlyRootFilesystem: true,
            runAsNonRoot: true,
          },
        }),
      ],
      podMetadata: {
        labels: label,
      },
      securityContext: {
        fsGroupChangePolicy: 'Always',
        runAsNonRoot: true,
      },
    });

    // Define the Service to expose the Deployment
    new k8s.Service(this, 'web-service', {
      metadata: { name: 'web-service' },
      type: k8s.ServiceType.LOAD_BALANCER,
      ports: [{ port: 80, targetPort: 80 }],
      selector: label,
    });
  }
}

const app = new App();
new MyChart(app, 'my-chart');
app.synth();
