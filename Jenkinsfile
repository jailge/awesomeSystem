pipeline {
    agent {
        kubernetes{
            label 'jenkins-slave'
        }

    }
    environment{
        IMAGE_NAME = "${params.image_name}"
        BUILD_NUMBER = "${params.BUILD_NUMBER}"
        DOCKER_USERNAME = credentials('DOCKER_USERNAME')
        DOCKER_PASSWORD = credentials('DOCKER_PASSWORD')
    }
    stages {
        stage('docker login') {
            steps{
                sh(script: """
                    docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
                """, returnStdout: true)
            }
        }

        stage('git clone') {
            steps{
                sh(script: """
                    git clone https://github.com/jailge/awesomeSystem.git
                """, returnStdout: true)
            }
        }

        stage('docker build') {
            steps{
                sh script: '''
                #!/bin/bash
                cd $WORKSPACE/awesomeSystem/
                docker build . --network host -t jailge/$IMAGE_NAME:$BUILD_NUMBER
                '''
            }
        }

        stage('docker push') {
            steps{
                sh(script: """
                    docker push jailge/$IMAGE_NAME:$BUILD_NUMBER
                """)
            }
        }

    }
}