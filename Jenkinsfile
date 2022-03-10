pipeline {
    agent none
    stages {
        stage('build and test') {
            agent { docker { image 'golang:1.13' } }
            environment {
                GOCACHE = '/tmp/gocache'
            }
            steps {
                sh 'go build'
                sh 'go test ./...'
            }
        }
        stage('deploy') {
            agent any
            environment {
                BUILD_NUMBER_BASE = '0'
                VERSION_MAJOR = '0'
                VERSION_MINOR = '1'
                VERSION_PATCH = "${env.BUILD_NUMBER.toInteger() - BUILD_NUMBER_BASE.toInteger()}"
            }
            when {
                branch 'master'
            }
            steps {
                script {
                    docker.withRegistry('', 'docker-hub') {
                        def customImage = docker.build("riyaz1994/gocode:$VERSION_MAJOR.$VERSION_MINOR.$VERSION_PATCH")
                        customImage.push()
                    }
                }
            }
        }
    }
}
