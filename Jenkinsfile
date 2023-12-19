pipeline {
    agent {
        label 'fire'
    }

    environment {
        WEBHOOK_URL_DISCORD = credentials('WEBHOOK_URL_DISCORD')
    }

    stages {
        stage('Test Goapps'){
            agent {
                docker {
                    image 'golang:1.21.4-alpine3.18'
                    label 'fire'
                }
            }

            steps {
                echo "Testing Golang Apps"
                sh 'GOCACHE=/tmp/ go test -v ./healthcheck'
                sh 'GOCACHE=/tmp/ go test -v ./helper'
            }   
        }

        stage('Build') {
            steps {
                echo "Building Apps"
                // sh 'docker build -t gcr.io/ancient-alloy-406700/goapps:${BUILD_NUMBER} .'
            }
        }

        stage('Push to DockerHub') {
            // environment {
            // }

            steps {
                echo "Pushing to DockerHub"
            }

            // post {
            //     success {
            //         echo "Post Success"
            //         discordSend description: "Jenkins Pipeline Push", footer: "Push Success image goapps:${BUILD_NUMBER}", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
            //     }

            //     failure {
            //         echo "Post Failure"
            //         discordSend description: "Jenkins Pipeline Push", footer: "Push Failure image goapps:${BUILD_NUMBER}", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
            //     }
            // }
        }

        stage('Deploy') {
            steps {
                echo "Deploying apps"
            }
        }

    }
    
    post {
        success {
            echo "Post Success"
            discordSend description: "Jenkins Pipeline Deploy", footer: "Fundhub Production Deployment Status: Success", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
        }
        failure {
            echo "Post Failure"
            discordSend description: "Jenkins Pipeline Deploy", footer: "Fundhub Production Deployment Status: Failure", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK_URL_DISCORD"
        }
    }
}