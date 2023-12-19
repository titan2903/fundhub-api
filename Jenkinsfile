pipeline {
    agent {
        label 'fire'
    }

    environment {
        WEBHOOK_URL_DISCORD = credentials('WEBHOOK_URL_DISCORD')
        DOCKER_USERNAME = credentials('DOCKER_USERNAME')
        DOCKER_PASSWORD = credentials('DOCKER_PASSWORD')
    }

    stages {
        stage('Test Fundhub'){
            agent {
                docker {
                    image 'golang:1.21.4-alpine3.18'
                    label 'fire'
                }
            }

            steps {
                echo "Testing Golang Apps"
                sh 'GOCACHE=/tmp/ go test -v ./...'
            }   
        }

        stage('Build') {
            steps {
                echo "Building Apps"
                sh 'docker build -t $DOCKER_USERNAME/fundhub:latest .'
            }
        }

        stage('Push to DockerHub') {
            steps {
                echo "Pushing to DockerHub"
                sh "docker login -u $DOCKER_USERNAME -p ${DOCKER_PASSWORD}"
                sh "docker push $DOCKER_USERNAME/fundhub-api:latest"
            }

            post {
                aborted {
                    echo "Post Aborted"
                    discordSend description: "Push Docker Image", footer: "Push image fundhub-api:latest to DockerHub Status: Aborted", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
                }
                success {
                    echo "Post Success"
                    discordSend description: "Push Docker Image", footer: "Push image fundhub-api:latest to DockerHub Status: Success", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
                }

                failure {
                    echo "Post Failure"
                    discordSend description: "Push Docker Image", footer: "Push image fundhub-api:latest to DockerHub Status: Failure", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
                }
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploying apps"
            }
        }

    }
    
    post {
        aborted {
            echo "Post Aborted"
            discordSend description: "Fundhub Deployment", footer: "Fundhub Production Deployment Status: Aborted", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
        }

        success {
            echo "Post Success"
            discordSend description: "Fundhub Deployment", footer: "Fundhub Production Deployment Status: Success", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
        }
        
        failure {
            echo "Post Failure"
            discordSend description: "Fundhub Deployment", footer: "Fundhub Production Deployment Status: Failure", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
        }
    }
}