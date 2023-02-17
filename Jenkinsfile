pipeline {
  agent {
    kubernetes {
        yaml """
        apiVersion: v1
        kind: Pod
        metadata:
          labels:
            jenkins: agent
            k8s-app: jenkins-agent
        spec:
          containers:
          - name: golang
            image: golang:1.19.3-bullseye
            command:
              - cat
            tty: true
        """
    }
  }
  stages {
    stage('Compile') {
      steps {
        container('golang') {
          sh 'go build'
        }
      }
    }

    stage('Test') {
      steps {
        container('golang') {
          sh 'go test ./...'
        }
      }
    }

    stage ('Release') {
      environment {
        GITHUB_TOKEN = credentials('github-sa')
      }

      steps {
        container('golang') {
          sh 'curl -sfL https://goreleaser.com/static/run | bash -s -- --snapshot'
        }
        archiveArtifacts allowEmptyArchive: true, artifacts: 'dist/*.tar.gz, dist/*.zip, dist/*.txt, dist/*.json', onlyIfSuccessful: true
      }
    }
  }
}
