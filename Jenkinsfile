pipeline {
    agent any

    environment {
        GO_VERSION = '1.23.4'  // Define the Go version you're using
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout the code from GitHub
                git credentialsId: 4fe29774-589c-44fd-adb0-66fa3efdeeac, url: https://github.com/mritunjaygour1/tm.git
            }
        }

        stage('Install Go') {
            steps {
                // Install Go (if not already installed on the Jenkins server)
                script {
                    if (!fileExists('/usr/local/go/bin/go')) {
                        sh 'wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz'
                        sh 'tar -C /usr/local -xvzf go$GO_VERSION.linux-amd64.tar.gz'
                    }
                }
            }
        }

        stage('Build') {
            steps {
                // Build the Go application
                sh '''
                    export PATH=$PATH:/usr/local/go/bin
                    go version
                    go build -o myapp .
                '''
            }
        }

        stage('Test') {
            steps {
                // Run Go tests
                sh '''
                    export PATH=$PATH:/usr/local/go/bin
                    go test ./...
                '''
            }
        }

        stage('Deploy') {
            steps {
                // Deploy the application (you can customize this to your own needs)
                echo 'Deploying the application'
                // Add deployment steps here, such as pushing to Docker, deploying to Kubernetes, etc.
            }
        }
    }

    post {
        always {
            cleanWs()  // Clean workspace after build
        }

        success {
            echo 'Build completed successfully!'
        }

        failure {
            echo 'Build failed!'
        }
    }
}
