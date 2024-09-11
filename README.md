# Instalando para rodar lambda localmente

## 1 - (if needed) Install the lambda runtime interface emulator

Instruções oficiais:

https://github.com/aws/aws-lambda-runtime-interface-emulator?tab=readme-ov-file#installing

```bash
# For x86_64 systems
mkdir -p ~/.aws-lambda-rie && curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie && chmod +x ~/.aws-lambda-rie/aws-lambda-rie

# For ARM64 systems
mkdir -p ~/.aws-lambda-rie && curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie-arm64 && chmod +x ~/.aws-lambda-rie/aws-lambda-rie
```

## Run locally:

- Depois disso só rodar `go run main.go`

# Deploy AWS Lambda

![Deploy part 1](https://github.com/viquitorreis/aws-serverless-echo-adapter/blob/main/aws-1.png?raw=true)

![Deploy part 2](https://github.com/viquitorreis/aws-serverless-echo-adapter/blob/main/aws-2.png?raw=true)

    Upload the binary (.zip - compressed) on "Upload"