## Go のセットアップ

https://go.dev/doc/install

## 開発環境でのサーバー立ち上げ

```bash
GO_ENV=development go run main.go
```

## 容量を空けるために開発環境に残っている Docker イメージを全削除する方法

```bash
docker rmi -f $(docker images -q)
```

## キャッシュとビルドデータを全削除する方法

```bash
docker builder prune
```

## ビルド

```bash
sam build
```

## デプロイ

```bash
sam deploy --guided --capabilities CAPABILITY_NAMED_IAM
```

The command will package and deploy your application to AWS, with a series of prompts:

- **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
- **AWS Region**: The AWS region you want to deploy your app to.
- **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
- **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modifies IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
- **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

## SAM 削除

```bash
sam delete <stack-name>
```

## 参考リンク

- [Go のプロジェクト構成](https://zenn.dev/nobonobo/articles/4fb018a24f9ee9)
- [DTO とは](https://zenn.dev/7oh/articles/6338a8ccd470c7#%E3%83%AA%E3%83%9D%E3%82%B8%E3%83%88%E3%83%AA%E3%81%AE%E4%BD%9C%E6%88%90)
- [Go Echo サーバーを SAM で公開する方法](https://zenn.dev/ryichk/articles/90d492d7874b1f#3.-sam%E3%81%AEtemplate.yaml%E3%82%92%E4%BD%9C%E6%88%90)
- [AWS Lambda Function URLs と Amazon API Gateway の違い](https://serverless.co.jp/blog/j94zz_4-m/)
- [API Gateway 挟むパターンでもパス全通しできる Lambda プロキシ統合](<https://zenn.dev/tkg216/articles/c9aede81c20a8e#%E3%83%97%E3%83%AD%E3%82%AD%E3%82%B7%E3%83%AA%E3%82%BD%E3%83%BC%E3%82%B9(%7Bproxy%2B%7D)%E3%81%AE%E6%9C%89%E7%84%A1%E3%81%A7%E6%AF%94%E8%BC%83%E8%BB%B8>)
