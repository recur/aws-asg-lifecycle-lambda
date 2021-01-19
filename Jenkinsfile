#!/usr/bin/env groovy

library 'defn/jenkins-kiki@main'

def config = [
  name: 'recur/aws-asg-lifecycle-lambda',
  roleId: 'c0f5f093-4083-91c2-246a-b78a1df298da',
  secretValues: [
    [vaultKey: 'MEH1'],
    [vaultKey: 'MEH2']
  ]
]

goreleaserMain(config)
