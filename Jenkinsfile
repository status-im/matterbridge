pipeline {
  agent {
    docker {
      label 'linuxcontainer'
      image 'harbor.status.im/infra/ci-build-containers:linux-base-1.0.0'
      args '--volume=/var/run/docker.sock:/var/run/docker.sock ' +
           '--user jenkins'
    }
  }

  parameters {
    string(
      name: 'GIT_REF',
      defaultValue: 'wakuv2',
      description: 'Branch, tag, or commit to build.'
    )
    string(
      name: 'IMAGE_NAME',
      description: 'Docker image name.',
      defaultValue: params.IMAGE_NAME ?: 'status-im/matterbridge',
    )
    string(
      name: 'IMAGE_TAG',
      description: 'Docker image tag.',
      defaultValue: params.IMAGE_TAG ?: ''
    )
    string(
      name: 'DOCKER_CRED',
      description: 'Name of Docker Registry credential.',
      defaultValue: params.DOCKER_CRED ?: 'harbor-status-im-robot',
    )
    string(
      name: 'DOCKER_REGISTRY_URL',
      description: 'URL of the Docker Registry',
      defaultValue: params.DOCKER_REGISTRY_URL ?: 'https://harbor.status.im'
    )
  }

  options {
    timestamps()
    disableRestartFromStage()
    buildDiscarder(logRotator(
      numToKeepStr: '10',
      daysToKeepStr: '30',
    ))
  }

  stages {
    stage('Checkout') {
      steps {
        checkout([
          $class: 'GitSCM',
          branches: [[name: params.GIT_REF]],
          userRemoteConfigs: scm.userRemoteConfigs
        ])
      }
    }
    stage('Build') {
      steps { script {
        env.IMAGE_TAG = params.IMAGE_TAG ?: env.GIT_COMMIT.take(8)

        image = docker.build(
          "${params.IMAGE_NAME}:${env.IMAGE_TAG}",
          "--build-arg='GIT_COMMIT=${GIT_COMMIT.take(8)}' ."
        )
      } }
    }

    stage('Push') {
      steps { script {
        withDockerRegistry([
          credentialsId: params.DOCKER_CRED,
          url: params.DOCKER_REGISTRY_URL
        ]) {
          image.push(env.IMAGE_TAG)
        }
      } }
    }
  }
}