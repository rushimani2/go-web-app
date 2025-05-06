"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const cdk8s_1 = require("cdk8s");
const k8s = require("cdk8s-plus-25");
class MyChart extends cdk8s_1.Chart {
    constructor(scope, id) {
        super(scope, id);
        // Create a Deployment resource
        new k8s.Deployment(this, 'web-deployment', {
            metadata: { name: 'web' },
            containers: [
                {
                    image: 'nginx', // The container image for the Deployment
                    name: 'nginx-container',
                }
            ]
        });
    }
}
const app = new cdk8s_1.App();
new MyChart(app, 'my-chart');
app.synth();
