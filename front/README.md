## Nodejs のセットアップ

https://asdf-vm.com/ja-jp/guide/getting-started.html

## サーバー立ち上げ

```bash
npm install
npm run dev
```

## ビルド

```bash
npm run build
```

## 開発環境で本番サーバー立ち上げ

```bash
npx serve@latest out
```

## 静的ホスティング用の S3 スタックデプロイ

```bash
aws cloudformation create-stack --stack-name junbanmachi-front-app --template-body file://infra/s3-static-hosting.yaml
```

## デプロイ(S3 にビルドファイル配置)

```bash
aws s3 sync ./out s3://junbanmachi-bucket
```

## 参考リンク

- [Nextjs を SAM で公開する方法](https://github.com/awslabs/aws-lambda-web-adapter/tree/main/examples/nextjs)
- [Static Export](https://nextjs.org/docs/pages/building-your-application/deploying/static-exports)
